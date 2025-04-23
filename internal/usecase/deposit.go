package usecase

import (
	"context"

	"github.com/shopspring/decimal"
	"imansohibul.my.id/account-domain-service/entity"
	"imansohibul.my.id/account-domain-service/util"
)

type depositUsecase struct {
	accountRepository     AccountRepository
	transactionRepository TransactionRepository
	transactionManager    TransactionManager
	logger                util.Logger
}

func NewDepositUsecase(
	accountRepository AccountRepository,
	transactionRepository TransactionRepository,
	transactionManager TransactionManager,
	logger util.Logger,
) *depositUsecase {
	return &depositUsecase{
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
		transactionManager:    transactionManager,
		logger:                logger,
	}
}

func (d depositUsecase) Deposit(ctx context.Context, accountNumber string, amount decimal.Decimal) (*entity.Transaction, error) {
	var (
		applyLock = true
		err       error
		logger    = d.logger.WithDuration(
			ctx,
			"depositUsecase.Deposit",
			map[string]interface{}{
				"account_number": accountNumber,
				"amount":         amount,
			},
		)
	)

	defer logger(&err)

	transaction := new(entity.Transaction)

	err = d.transactionManager.WithTransaction(ctx, func(ctx context.Context) error {
		// Find account by account number and lock it for update
		// to prevent concurrent access and update the balance
		account, err := d.accountRepository.FindByAccountNumber(ctx, entity.AccountTypeSaving, accountNumber, applyLock)
		if err != nil {
			return err
		}

		transaction.AccountID = account.ID
		transaction.Amount = amount
		transaction.Type = entity.TransactionTypeCredit
		transaction.InitialBalance = account.Balance
		transaction.FinalBalance = account.Balance.Add(amount)
		transaction.Currency = account.Currency

		account.Balance = account.Balance.Add(amount)
		_, err = d.accountRepository.UpdateAccount(ctx, account)
		if err != nil {
			return err
		}

		transaction, err = d.transactionRepository.CreateTransaction(ctx, transaction)
		if err != nil {
			return err
		}

		return nil

	})

	return transaction, err
}
