package usecase

import (
	"context"

	"github.com/shopspring/decimal"
	"imansohibul.my.id/account-domain-service/entity"
	"imansohibul.my.id/account-domain-service/util"
)

type withdrawUsecase struct {
	accountRepository     AccountRepository
	transactionRepository TransactionRepository
	transactionManager    TransactionManager
	logger                util.Logger
}

func NewWithdrawUsecase(
	accountRepository AccountRepository,
	transactionRepository TransactionRepository,
	transactionManager TransactionManager,
	logger util.Logger,
) *withdrawUsecase {
	return &withdrawUsecase{
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
		transactionManager:    transactionManager,
		logger:                logger,
	}
}

func (w withdrawUsecase) Withdraw(ctx context.Context, accountNumber string, amount decimal.Decimal) (*entity.Transaction, error) {
	var (
		applyLock = true
		err       error
		logger    = w.logger.WithDuration(
			ctx,
			"withdrawUsecase.Withdraw",
			map[string]interface{}{
				"account_number": accountNumber,
				"amount":         amount,
			},
		)
	)

	defer logger(&err)

	transaction := new(entity.Transaction)

	err = w.transactionManager.WithTransaction(ctx, func(ctx context.Context) error {
		// Find account by account number and lock it for update
		// to prevent concurrent access and update the balance
		account, err := w.accountRepository.FindByAccountNumber(ctx, entity.AccountTypeSaving, accountNumber, applyLock)
		if err != nil {
			return err
		}

		if account.Balance.LessThan(amount) {
			return entity.ErrInsufficientBalance
		}

		transaction.AccountID = account.ID
		transaction.Amount = amount
		transaction.Type = entity.TransactionTypeDebit
		transaction.InitialBalance = account.Balance
		transaction.FinalBalance = account.Balance.Sub(amount)
		transaction.Currency = account.Currency

		transaction, err = w.transactionRepository.CreateTransaction(ctx, transaction)
		if err != nil {
			return err
		}

		account.Balance = account.Balance.Sub(amount)
		if _, err := w.accountRepository.UpdateAccount(ctx, account); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
