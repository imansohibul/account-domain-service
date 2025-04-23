package usecase

import (
	"context"

	"github.com/shopspring/decimal"
	"imansohibul.my.id/account-domain-service/entity"
	"imansohibul.my.id/account-domain-service/util"
)

type getBalanceUsecase struct {
	accountRepository AccountRepository
	logger            util.Logger
}

func NewGetBalanceUsecase(
	accountRepository AccountRepository,
	logger util.Logger,
) *getBalanceUsecase {
	return &getBalanceUsecase{
		accountRepository: accountRepository,
		logger:            logger,
	}
}

func (g getBalanceUsecase) GetBalance(ctx context.Context, accountNumber string) (decimal.Decimal, error) {
	var (
		err       error
		applyLock = false
		logger    = g.logger.WithDuration(
			ctx,
			"getBalanceUsecase.GetBalance",
			map[string]interface{}{
				"account_number": accountNumber,
			},
		)
	)

	defer logger(&err)

	account, err := g.accountRepository.FindByAccountNumber(ctx, entity.AccountTypeSaving, accountNumber, applyLock)
	if err != nil {
		return decimal.Zero, err
	}

	return account.Balance, nil
}
