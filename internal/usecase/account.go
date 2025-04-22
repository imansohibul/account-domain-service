package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"imansohibul.my.id/account-domain-service/entity"
	"imansohibul.my.id/account-domain-service/util"
)

// DefaultAccountNumberLength is the length of the account number
const DefaultAccountNumberLength = 10

// DefaultMaxRetries is the maximum number of retries for creating an account
const DefaultMaxRetries = 3

// accountUsecase implements the AccountUsecase interface
type accountUsecase struct {
	accountRepository          AccountRepository
	transactionManager         TransactionManager
	customerRepository         CustomerRepository
	customerIdentityRepository CustomerIdentityRepository
	transactionRepository      TransactionRepository
}

func NewAccountUsecase(
	accountRepository AccountRepository,
	transactionManager TransactionManager,
	customerRepository CustomerRepository,
	customerIdentityRepository CustomerIdentityRepository,
	transactionRepository TransactionRepository,
) *accountUsecase {
	return &accountUsecase{
		accountRepository:          accountRepository,
		transactionManager:         transactionManager,
		customerRepository:         customerRepository,
		customerIdentityRepository: customerIdentityRepository,
		transactionRepository:      transactionRepository,
	}
}

func (a accountUsecase) CreateAccount(ctx context.Context, params *entity.CreateAccountParams) (*entity.Account, error) {
	// Check if phone number already exists
	customer, err := a.customerRepository.FindByPhoneNumber(ctx, params.PhoneNumber)
	if err != nil && err != entity.ErrNotFound {
		return nil, err
	}

	if customer != nil {
		return nil, entity.ErrPhoneNumberAlreadyExists
	}

	// Check if customer already exists
	customerIdentity, err := a.customerIdentityRepository.FindByIdentity(ctx, entity.IdentityTypeNIK, params.IdentityNumber)
	if err != nil && err != entity.ErrNotFound {
		return nil, err
	}

	if customerIdentity != nil {
		return nil, entity.ErrCustomerAlreadyExists
	}

	account := new(entity.Account)

	err = a.transactionManager.WithTransaction(ctx, func(ctx context.Context) error {
		// Create customer
		customer, err := a.customerRepository.CreateCustomer(ctx, &entity.Customer{
			Fullname:    params.Fullname,
			PhoneNumber: params.PhoneNumber,
		})
		if err != nil {
			return err
		}

		// Create customer identity
		customerIdentity, err = a.customerIdentityRepository.CreateCustomerIdentity(ctx, &entity.CustomerIdentity{
			CustomerID:     customer.ID,
			IdentityType:   entity.IdentityTypeNIK,
			IdentityNumber: params.IdentityNumber,
		})
		if err != nil {
			return err
		}

		account, err = a.createAccountWithRetry(ctx, customer, DefaultMaxRetries)
		if err != nil {
			return err
		}

		return nil
	})

	return account, err
}

// createAccountWithRetry validates if the account number is unique during the insert operation
func (a accountUsecase) createAccountWithRetry(ctx context.Context, customer *entity.Customer, maxRetries int) (*entity.Account, error) {
	var (
		err     error
		account = &entity.Account{
			CustomerID:  customer.ID,
			AccountType: entity.AccountTypeSaving,
			Status:      entity.AccountStatusActive,
			Balance:     decimal.Zero,
			Currency:    entity.CurrencyIDR,
		}
	)

	for attempt := 1; attempt <= maxRetries; attempt++ {
		account.AccountNumber, err = util.GenerateSecureNumber(DefaultAccountNumberLength)
		if err != nil {
			return nil, fmt.Errorf("failed to generate account number: %w", err)
		}

		// Attempt to insert the account into the database
		account, err = a.accountRepository.CreateAccount(ctx, account)
		if err != nil && err != entity.ErrDuplicateAccountNumber {
			return nil, fmt.Errorf("failed to create account: %w", err)
		} else if err != nil {
			fmt.Printf("Attempt %d: Duplicate account number %s, retrying...\n", attempt, account.AccountNumber)
			time.Sleep(time.Duration(attempt) * time.Second)
			continue
		}

		// Account successfully created, exit
		return account, nil
	}

	// If all attempts failed, return the last error encountered
	return nil, fmt.Errorf("failed to create account after %d attempts: %w", maxRetries, err)
}
