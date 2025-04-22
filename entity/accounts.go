package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

// AccountType represents the type of account
type AccountType int16

// AccountType is an enumeration of account types
// The enumeration values are:
// 0 - Unspecified
// 1 - Saving
const (
	AccountTypeUnspecified AccountType = iota
	AccountTypeSaving
)

// AccountStatus represents the status of an account
type AccountStatus int16

// AccountStatus is an enumeration of account statuses
// The enumeration values are:
// 0 - Unspecified
// 1 - Active
const (
	AccountStatusUnspecified AccountStatus = iota
	AccountStatusActive
)

type Account struct {
	ID            uint
	CustomerID    uint
	AccountType   AccountType
	AccountNumber string
	Balance       decimal.Decimal
	Currency      Currency
	Status        AccountStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// CreateAccountParams represents the request to create an account
// Will be used as parameters for the use case of creating an account
type CreateAccountParams struct {
	Fullname       string
	PhoneNumber    string
	IdentityNumber string
}
