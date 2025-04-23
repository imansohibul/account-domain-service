package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"imansohibul.my.id/account-domain-service/entity"
)

type accountHandler struct {
	createAccountUsecase CreateAccountUsecase
	depositUsecaase      DepositUsecase
	withdrawUsecase      WithdrawUsecase
	getBalanceUsecase    GetBalanceUsecase
}

func NewAccountHandler(
	createAccountUsecase CreateAccountUsecase,
	depositUsecase DepositUsecase,
	withdrawUsecase WithdrawUsecase,
	getBalanceUsecase GetBalanceUsecase,
) *accountHandler {
	return &accountHandler{
		createAccountUsecase: createAccountUsecase,
		depositUsecaase:      depositUsecase,
		withdrawUsecase:      withdrawUsecase,
		getBalanceUsecase:    getBalanceUsecase,
	}
}

func (a accountHandler) CreateAccount(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req = new(CreateAccountRequest)
	)

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

	account, err := a.createAccountUsecase.CreateAccount(ctx, params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	return c.JSON(http.StatusOK, &CreateAccountResponse{
		AccountNumber: account.AccountNumber,
	})
}

func (a accountHandler) Deposit(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req = new(DepositRequest)
	)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request"})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	transaction, err := a.depositUsecaase.Deposit(ctx, req.AccountNumber, req.GetAmount())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	return c.JSON(http.StatusOK, &DepositResponse{
		AccountBalance: transaction.FinalBalance,
	})
}

func (a accountHandler) Withdraw(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req = new(WithdrawRequest)
	)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request"})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	transaction, err := a.withdrawUsecase.Withdraw(ctx, req.AccountNumber, req.GetAmount())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	return c.JSON(http.StatusOK, &DepositResponse{
		AccountBalance: transaction.FinalBalance,
	})
}

func (a accountHandler) GetBalance(c echo.Context) error {
	var (
		ctx           = c.Request().Context()
		accountNumber = c.Param("account_number") // Get account number from path parameter
	)

	// If account number is not provided, return bad request
	if accountNumber == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "Invalid request"})
	}

	balance, err := a.getBalanceUsecase.GetBalance(ctx, accountNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	return c.JSON(http.StatusOK, &GetBalanceResponse{
		AccountBalance: balance,
	})
}
