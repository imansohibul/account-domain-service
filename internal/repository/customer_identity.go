package repository

import (
	"context"
	"errors"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"imansohibul.my.id/account-domain-service/entity"
)

type customerIdentityRepository struct {
	db rel.Repository
}

type customerIdentity struct {
	ID             uint      `db:"id"`
	CustomerID     uint      `db:"customer_id"`
	IdentityType   int       `db:"identity_type"`
	IdentityNumber string    `db:"identity_number"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func NewCustomerIdentityRepository(db rel.Repository) *customerIdentityRepository {
	return &customerIdentityRepository{db: db}
}

func (c customerIdentityRepository) CreateCustomerIdentity(ctx context.Context, customerIdentity *entity.CustomerIdentity) (*entity.CustomerIdentity, error) {
	customerIdentityRecord := c.fromEntityCustomerIdentity(customerIdentity)
	err := c.db.Insert(ctx, customerIdentityRecord)
	if err != nil && !errors.Is(err, rel.ErrUniqueConstraint) {
		return nil, err
	} else if errors.Is(err, rel.ErrUniqueConstraint) {
		return nil, entity.ErrCustomerIdentityAlreadyExists
	}

	return c.toEntityCustomerIdentity(customerIdentityRecord), nil
}

func (c customerIdentityRepository) FindByIdentity(ctx context.Context, identityType entity.CustomerIdentityType, identityNumber string) (*entity.CustomerIdentity, error) {
	customerIdentityRecord := new(customerIdentity)
	err := c.db.Find(ctx, customerIdentityRecord, where.Eq("identity_type = ?", int(identityType)), where.Eq("identity_number = ?", identityNumber))
	if err != nil && errors.Is(err, rel.ErrNotFound) {
		return nil, entity.ErrCustomerIdentityNotFound
	} else if err != nil {
		return nil, err
	}

	return c.toEntityCustomerIdentity(customerIdentityRecord), nil
}

func (c customerIdentityRepository) fromEntityCustomerIdentity(customerIdentityRecord *entity.CustomerIdentity) *customerIdentity {
	return &customerIdentity{
		ID:             customerIdentityRecord.ID,
		CustomerID:     customerIdentityRecord.CustomerID,
		IdentityType:   int(customerIdentityRecord.IdentityType),
		IdentityNumber: customerIdentityRecord.IdentityNumber,
	}
}
func (c customerIdentityRepository) toEntityCustomerIdentity(customerIdentityRecord *customerIdentity) *entity.CustomerIdentity {
	return &entity.CustomerIdentity{
		ID:             customerIdentityRecord.ID,
		CustomerID:     customerIdentityRecord.CustomerID,
		IdentityType:   entity.CustomerIdentityType(customerIdentityRecord.IdentityType),
		IdentityNumber: customerIdentityRecord.IdentityNumber,
	}
}
