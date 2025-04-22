package entity

import "time"

// CustomerIdentityType represents the type of customer identity
// CustomerIdentityType is an enumeration of customer identity types
// The enumeration values are:
// 0 - Unspecified
// 1 - NIK (Nomor Induk Kependudukan)
type CustomerIdentityType int16

// Enumeration of customer identity types
const (
	IdentityTypeUnspecifed CustomerIdentityType = iota
	IdentityTypeNIK
)

// CustomerIdentity represents the identity of a customer
type CustomerIdentity struct {
	ID             uint
	CustomerID     uint
	IdentityType   CustomerIdentityType
	IdentityNumber string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
