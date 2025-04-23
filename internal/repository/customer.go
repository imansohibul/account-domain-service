package repository

import (
	"context"
	"errors"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"imansohibul.my.id/account-domain-service/entity"
)

type customerRepository struct {
	db rel.Repository
}

func NewCustomerRepository(db rel.Repository) *customerRepository {
	return &customerRepository{db: db}
}

type customer struct {
	ID          uint      `db:"id"`
	Fullname    string    `db:"fullname"`
	PhoneNumber string    `db:"phone_number"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (c customerRepository) CreateCustomer(ctx context.Context, newCustomer *entity.Customer) (*entity.Customer, error) {
	customerRecord := c.fromEntityCustomer(newCustomer)

	err := c.db.Insert(ctx, customerRecord)
	if err != nil && !errors.Is(err, rel.ErrUniqueConstraint) {
		return nil, err
	} else if errors.Is(err, rel.ErrUniqueConstraint) {
		return nil, entity.ErrPhoneNumberAlreadyExists
	}

	return c.toEntityCustomer(customerRecord), nil
}

func (c customerRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.Customer, error) {
	customerRecord := new(customer)
	err := c.db.Find(ctx, customerRecord, where.Eq("phone_number", phoneNumber))
	if err != nil && errors.Is(err, rel.ErrNotFound) {
		return nil, entity.ErrCustomerNotFound
	} else if err != nil {
		return nil, err
	}

	return c.toEntityCustomer(customerRecord), nil
}

func (c customerRepository) fromEntityCustomer(customerRecord *entity.Customer) *customer {
	return &customer{
		ID:          customerRecord.ID,
		Fullname:    customerRecord.Fullname,
		PhoneNumber: customerRecord.PhoneNumber,
	}
}

func (c customerRepository) toEntityCustomer(customerRecord *customer) *entity.Customer {
	return &entity.Customer{
		ID:          customerRecord.ID,
		Fullname:    customerRecord.Fullname,
		PhoneNumber: customerRecord.PhoneNumber,
	}
}
