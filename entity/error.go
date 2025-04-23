package entity

import "errors"

var (
	// ErrAccountNotFound is returned when the requested account is not found
	ErrAccountNotFound = errors.New("Nomor rekening tidak ditemukan")

	// ErrDuplicateAccountNumber is returned when an account with the same account number already exists
	ErrDuplicateAccountNumber = errors.New("Nomor rekening sudah terdaftar")

	// ErrCustomerNotFound is returned when the requested customer is not found
	ErrCustomerNotFound = errors.New("Nasabah tidak ditemukan")

	// ErrPhoneNumberAlreadyExists is returned when a customer with the same phone number already exists
	ErrPhoneNumberAlreadyExists = errors.New("Nomor telepon sudah terdaftar")

	// ErrCustomerIdentityNotFound is returned when the requested customer identity is not found
	ErrCustomerIdentityNotFound = errors.New("Identitas nasabah tidak ditemukan")

	//ErrCustomerIdentityAlreadyExists is returned when a customer with the same NIK already exists
	ErrCustomerIdentityAlreadyExists = errors.New("NIK sudah terdaftar")

	// ErrInvalidRequest is returned when the request is invalid
	ErrInvalidRequest = errors.New("Permintaan tidak valid")
)
