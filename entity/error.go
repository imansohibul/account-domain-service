package entity

import "errors"

var (
	// ErrNotFound is returned when the requested resource is not found
	ErrNotFound = errors.New("not found")

	// ErrCustomerAlreadyExists is returned when a customer with the same NIK already exists
	ErrCustomerAlreadyExists = errors.New("NIK sudah terdaftar")

	// ErrPhoneNumberAlreadyExists is returned when a customer with the same phone number already exists
	ErrPhoneNumberAlreadyExists = errors.New("Nomor telepon sudah terdaftar")

	// ErrDuplicateAccountNumber is returned when an account with the same account number already exists
	ErrDuplicateAccountNumber = errors.New("Nomor rekening sudah terdaftar")
)
