package entity

import (
    "time"

    "github.com/shopspring/decimal"
)

// TransactionType represents the type of transaction
type TransactionType int16

// TransactionType is an enumeration of transaction types
// The enumeration values are:		
// 0 - Unspecified
// 1 - Credit
// 2 - Debit
const (
		TransactionTypeUnspecified TransactionType = iota
		TransactionTypeCredit
		TransactionTypeDebit
)

type Transaction struct {
    ID             uint
    AccountID      uint
    Type           TransactionType
    Amount         decimal.Decimal
    InitialBalance decimal.Decimal
    FinalBalance   decimal.Decimal
    Currency       Currency
    CreatedAt      time.Time
    UpdatedAt      time.Time
}
