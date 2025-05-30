package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/avast/retry-go"
	"github.com/shopspring/decimal"
	"imansohibul.my.id/account-domain-service/entity"
	"imansohibul.my.id/account-domain-service/util"
)

// DefaultAccountNumberLength is the length of the account number
const DefaultAccountNumberLength = 10

// DefaultMaxRetries is the maximum number of retries for creating an account
const DefaultMaxRetries = 3

// createAccountUsecase implements the CreateAccountUsecase interface
type createAccountUsecase struct {
	accountRepository          AccountRepository
	transactionManager         TransactionManager
	customerRepository         CustomerRepository
	customerIdentityRepository CustomerIdentityRepository
	transactionRepository      TransactionRepository
	logger                     util.Logger
}

func NewCreateAccountUsecase(
	accountRepository AccountRepository,
	transactionManager TransactionManager,
	customerRepository CustomerRepository,
	customerIdentityRepository CustomerIdentityRepository,
	transactionRepository TransactionRepository,
	logger util.Logger,
) *createAccountUsecase {
	return &createAccountUsecase{
		accountRepository:          accountRepository,
		transactionManager:         transactionManager,
		customerRepository:         customerRepository,
		customerIdentityRepository: customerIdentityRepository,
		transactionRepository:      transactionRepository,
		logger:                     logger,
	}
}

func (a createAccountUsecase) CreateAccount(ctx context.Context, params *entity.CreateAccountParams) (*entity.Account, error) {
	var (
		err    error
		logger = a.logger.WithDuration(
			ctx,
			"createAccountUsecase.CreateAccount",
			map[string]interface{}{
				"fullname":        params.Fullname,
				"phone_number":    params.PhoneNumber,
				"identity_number": params.IdentityNumber,
			},
		)
	)

	defer logger(&err)

	// Validate phone number
	err = a.validatePhoneNumber(ctx, params.PhoneNumber)
	if err != nil {
		return nil, err
	}

	// Validate identity number
	err = a.validateIdentityNumber(ctx, params.IdentityNumber)
	if err != nil {
		return nil, err
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
		_, err = a.customerIdentityRepository.CreateCustomerIdentity(ctx, &entity.CustomerIdentity{
			CustomerID:     customer.ID,
			IdentityType:   entity.IdentityTypeNIK,
			IdentityNumber: params.IdentityNumber,
		})
		if err != nil {
			return err
		}

		// Check account
		account, err = a.createAccountWithRetry(ctx, customer, DefaultMaxRetries)
		if err != nil {
			return err
		}

		return nil
	})

	return account, err
}

// validatePhoneNumber checks if the phone number already exists
func (a createAccountUsecase) validatePhoneNumber(ctx context.Context, phoneNumber string) error {
	customer, err := a.customerRepository.FindByPhoneNumber(ctx, phoneNumber)
	if err != nil && err != entity.ErrCustomerNotFound {
		return err
	}

	if customer != nil {
		return entity.ErrPhoneNumberAlreadyExists
	}

	return nil
}

// validateIdentityNumber checks if the identity number already exists
func (a createAccountUsecase) validateIdentityNumber(ctx context.Context, identityNumber string) error {
	customerIdentity, err := a.customerIdentityRepository.FindByIdentity(ctx, entity.IdentityTypeNIK, identityNumber)
	if err != nil && err != entity.ErrCustomerIdentityNotFound {
		return err
	}

	if customerIdentity != nil {
		return entity.ErrCustomerIdentityAlreadyExists
	}

	return nil
}

// createAccountWithRetry validates if the account number is unique during the insert operation
func (a createAccountUsecase) createAccountWithRetry(ctx context.Context, customer *entity.Customer, maxRetries int) (*entity.Account, error) {
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

	// Retry mechanism using retry-go
	err = retry.Do(
		func() error {
			account.AccountNumber, err = util.GenerateSecureNumber(DefaultAccountNumberLength)
			if err != nil {
				return fmt.Errorf("failed to generate account number: %w", err)
			}

			account, err = a.accountRepository.CreateAccount(ctx, account)
			return err
		},
		retry.Attempts(uint(maxRetries)),
		retry.DelayType(retry.FixedDelay),
		retry.Delay(10*time.Microsecond),
		retry.RetryIf(func(err error) bool {
			return errors.Is(err, entity.ErrAccountAlreadyExists)
		}),
		retry.OnRetry(func(n uint, err error) {
			// log the retry attempt
			a.logger.Warn(ctx,
				"Retrying account creation due to duplicate account number",
				map[string]interface{}{
					"attempt": n,
					"error":   err.Error(),
				},
			)
		}),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create account after %d attempts: %w", maxRetries, err)
	}

	return account, nil
}
