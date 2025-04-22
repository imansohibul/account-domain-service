package accounts

import "github.com/labstack/echo/v4"

type accountHandler struct {
}

func NewAccountHandler() *accountHandler {
	return &accountHandler{}
}

func (a accountHandler) CreateAccount(c echo.Context) error {
	// TODO: Implement the logic to create an account
	return c.String(200, "Create Account")
}

func (a accountHandler) GetBalance(c echo.Context) error {
	// TODO: Implement the logic to get the balance of an account
	return c.String(200, "Get Account")
}

func (a accountHandler) Deposit(c echo.Context) error {
	// TODO: Implement the logic to deposit money into an account
	return c.String(200, "Deposit Account")
}

func (a accountHandler) Withdraw(c echo.Context) error {
	// TODO: Implement the logic to withdraw money from an account
	return c.String(200, "Withdraw Account")
}
