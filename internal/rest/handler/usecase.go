package handler

import (
	"context"

	"github.com/shopspring/decimal"
	"imansohibul.my.id/account-domain-service/entity"
)

//go:generate mockgen -destination=mock/usecase.go -package=mock -source=usecase.go

type CreateAccountUsecase interface {
	// CreateAccount creates a new account
	// returns the created account
	// returns an error if the account already exists or if the account creation fails
	CreateAccount(ctx context.Context, params *entity.CreateAccountParams) (*entity.Account, error)
}

type GetBalanceUsecase interface {
	// GetBalance retrieves the balance of an account
	// returns the balance of the account
	// returns an error if the account is not found or if the balance retrieval fails
	GetBalance(ctx context.Context, accountNumber string) (decimal.Decimal, error)
}

type DepositUsecase interface {
	// Deposit deposits money into an account
	// returns the transaction details of the deposit
	// returns an error if the account is not found or if the deposit fails
	Deposit(ctx context.Context, accountNumber string, amount decimal.Decimal) (*entity.Transaction, error)
}

type WithdrawUsecase interface {
	// Withdraw withdraws money from an account
	// returns the transaction details of the withdrawal
	// returns an error if the account is not found or if the withdrawal fails
	Withdraw(ctx context.Context, accountNumber string, amount decimal.Decimal) (*entity.Transaction, error)
}
