package entity

// DomainError represents a custom error with a Code and Message
type DomainError struct {
	Code    string
	Message string
}

func (e *DomainError) Error() string {
	return e.Message
}

// DomainError creates a new AppError instance
func NewDomainError(code, message string) *DomainError {
	return &DomainError{Code: code, Message: message}
}

var (
	// Account-related errors
	ErrAccountNotFound      = NewDomainError("ACCOUNT_NOT_FOUND", "Nomor rekening tidak ditemukan")
	ErrAccountAlreadyExists = NewDomainError("ACCOUNT_ALREADY_EXISTS", "Nomor rekening sudah terdaftar")
	ErrInsufficientBalance  = NewDomainError("ACCOUNT_INSUFFICIENT_BALANCE", "Saldo tidak mencukupi")

	// Customer-related errors
	ErrCustomerNotFound         = NewDomainError("CUSTOMER_NOT_FOUND", "Nasabah tidak ditemukan")
	ErrPhoneNumberAlreadyExists = NewDomainError("CUSTOMER_PHONE_NUMBER_EXISTS", "Nomor telepon sudah terdaftar")

	// Identity-related errors
	ErrCustomerIdentityNotFound      = NewDomainError("CUSTOMTER_IDENTITY_NOT_FOUND", "Identitas nasabah tidak ditemukan")
	ErrCustomerIdentityAlreadyExists = NewDomainError("CUSTOMER_IDENTITY_ALREADY_EXISTS", "NIK sudah terdaftar")

	// General errors
	ErrInvalidRequest = NewDomainError("INVALID_REQUEST", "Permintaan tidak valid")
)
