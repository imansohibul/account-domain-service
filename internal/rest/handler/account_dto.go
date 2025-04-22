package handler

import "github.com/shopspring/decimal"

// CreateAccountRequest	is the request body for creating an account
type CreateAccountRequest struct {
	Fullname       string `json:"nama" validate:"required,fullname"`
	PhoneNumber    string `json:"no_hp" validate:"required,e164"`
	IdentityNumber string `json:"nik" validate:"required,nik"`
}

// CreateAccountResponse is the response body for creating an account
type CreateAccountResponse struct {
	AccountNumber string `json:"no_rekening"`
}

// DepositRequest is the request body for depositing money into an account
type DepositRequest struct {
	AccountNumber string `json:"no_rekening" validate:"required"`
	Amount        int64  `json:"nominal" validate:"required, gt=0, lt=1000000000"`
}

// GetAmount converts the amount from int64 to decimal.Decimal
func (d DepositRequest) GetAmount() decimal.Decimal {
	return decimal.NewFromInt(d.Amount)
}

// DepositResponse is the response body for depositing money into an account
type DepositResponse struct {
	AccountBalance decimal.Decimal `json:"saldo"`
}

// WithdrawRequest is the request body for withdrawing money from an account
type WithdrawRequest struct {
	AccountNumber string `json:"no_rekening" validate:"required"`
	Amount        int64  `json:"nominal" validate:"required, gt=0, lt=100000000"`
}

// GetAmount converts the amount from int64 to decimal.Decimal
func (w WithdrawRequest) GetAmount() decimal.Decimal {
	return decimal.NewFromInt(w.Amount)
}

// WithdrawResponse is the response body for withdrawing money from an account
type WithdrawResponse struct {
	AccountBalance decimal.Decimal `json:"saldo"`
}

// GetBalanceResponse is the response body for getting the balance of an account
type GetBalanceResponse struct {
	AccountBalance decimal.Decimal `json:"saldo"`
}
