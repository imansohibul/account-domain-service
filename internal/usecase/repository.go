package usecase

import (
	"context"

	"imansohibul.my.id/account-domain-service/entity"
)

type TransactionManager interface {
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type AccountRepository interface {
	FindByAccountNumber(ctx context.Context, accountNumber string) (*entity.Account, error)
	CreateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error)
}

type CustomerRepository interface {
	CreateCustomer(ctx context.Context, customer *entity.Customer) (*entity.Customer, error)
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.Customer, error)
}

type CustomerIdentityRepository interface {
	CreateCustomerIdentity(ctx context.Context, customerIdentity *entity.CustomerIdentity) (*entity.CustomerIdentity, error)
	FindByIdentity(ctx context.Context, identityType entity.CustomerIdentityType, identityNumber string) (*entity.CustomerIdentity, error)
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
}
