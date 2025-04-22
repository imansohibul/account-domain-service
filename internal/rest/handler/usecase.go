package accounts

import (
	"github.com/shopspring/decimal"
	"imansohibul.my.id/account-domain-service/entity"
)

// AccountUsecase defines the interface for account use cases
type AccountUsecase interface {
	// CreateAccount creates a new account
	CreateAccount(params *entity.CreateAccountParams) (*entity.Account, error)

	// GetBalance retrieves the balance of an account
	GetBalance(accountNumber string) (decimal.Decimal, error)

	// Deposit deposits money into an account
	Deposit(accountNumber string, amount decimal.Decimal) error

	// Withdraw withdraws money from an account
	Withdraw(accountNumber string, amount decimal.Decimal) error
}
