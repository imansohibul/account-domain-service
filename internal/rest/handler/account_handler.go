package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"imansohibul.my.id/account-domain-service/entity"
)

type accountHandler struct {
	accountUsecase AccountUsecase
}

func NewAccountHandler(accountUsecase AccountUsecase) *accountHandler {
	return &accountHandler{
		accountUsecase: accountUsecase,
	}
}

func (a accountHandler) CreateAccount(c echo.Context) error {
	req := new(CreateAccountRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request"})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	params := &entity.CreateAccountParams{
		Fullname:       req.Fullname,
		PhoneNumber:    req.PhoneNumber,
		IdentityNumber: req.IdentityNumber,
	}

	account, err := a.accountUsecase.CreateAccount(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	return c.JSON(http.StatusOK, &CreateAccountResponse{
		AccountNumber: account.AccountNumber,
	})
}

func (a accountHandler) Deposit(c echo.Context) error {
	// TODO: Implement the logic to deposit money into an account
	return c.String(200, "Deposit Account")
}

func (a accountHandler) Withdraw(c echo.Context) error {
	// TODO: Implement the logic to withdraw money from an account
	return c.String(200, "Withdraw Account")
}

func (a accountHandler) GetBalance(c echo.Context) error {
	// TODO: Implement the logic to get the balance of an account
	return c.String(200, "Get Account")
}
