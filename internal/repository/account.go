package repository

import (
	"context"
	"errors"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/shopspring/decimal"
	"imansohibul.my.id/account-domain-service/entity"
)

type accountRepository struct {
	db rel.Repository
}

type account struct {
	ID            uint            `db:"id"`
	CustomerID    uint            `db:"customer_id"`
	AccountType   int             `db:"account_type"`
	AccountNumber string          `db:"account_number"`
	Balance       decimal.Decimal `db:"balance"`
	Currency      int             `db:"currency"`
	Status        int             `db:"status"`
	CreatedAt     time.Time       `db:"created_at"`
	UpdatedAt     time.Time       `db:"updated_at"`
}

func NewAccountRepository(db rel.Repository) *accountRepository {
	return &accountRepository{db: db}
}

func (a accountRepository) CreateAccount(ctx context.Context, newAccount *entity.Account) (*entity.Account, error) {
	accountRecord := a.fromEntityAccount(newAccount)

	err := a.db.Insert(ctx, accountRecord)
	if err != nil && !errors.Is(err, rel.ErrUniqueConstraint) {
		return nil, err
	} else if errors.Is(err, rel.ErrUniqueConstraint) {
		return nil, entity.ErrAccountAlreadyExists
	}

	return a.toEntityAccount(accountRecord), nil
}

func (a accountRepository) FindByAccountNumber(ctx context.Context, accountType entity.AccountType, accountNumber string, lock bool) (*entity.Account, error) {
	querier := []rel.Querier{
		where.Eq("account_type", int(accountType)),
		where.Eq("account_number", accountNumber),
	}

	if lock {
		querier = append(querier, rel.ForUpdate())
	}

	accountRecord := new(account)
	err := a.db.Find(ctx, accountRecord, querier...)
	if err != nil && errors.Is(err, rel.ErrNotFound) {
		return nil, entity.ErrAccountNotFound
	} else if err != nil {
		return nil, err
	}

	return a.toEntityAccount(accountRecord), nil
}

func (a accountRepository) UpdateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	accountRecord := a.fromEntityAccount(account)
	err := a.db.Update(ctx, accountRecord)
	if err != nil {
		return nil, err
	}

	return a.toEntityAccount(accountRecord), nil
}

func (a accountRepository) fromEntityAccount(accountEntity *entity.Account) *account {
	return &account{
		ID:            accountEntity.ID,
		CustomerID:    accountEntity.CustomerID,
		AccountType:   int(accountEntity.AccountType),
		AccountNumber: accountEntity.AccountNumber,
		Balance:       accountEntity.Balance,
		Currency:      int(accountEntity.Currency),
		Status:        int(accountEntity.Status),
		CreatedAt:     accountEntity.CreatedAt,
		UpdatedAt:     accountEntity.UpdatedAt,
	}
}

func (a accountRepository) toEntityAccount(accountRecord *account) *entity.Account {
	return &entity.Account{
		ID:            accountRecord.ID,
		CustomerID:    accountRecord.CustomerID,
		AccountType:   entity.AccountType(accountRecord.AccountType),
		AccountNumber: accountRecord.AccountNumber,
		Balance:       accountRecord.Balance,
		Currency:      entity.Currency(accountRecord.Currency),
		Status:        entity.AccountStatus(accountRecord.Status),
		CreatedAt:     accountRecord.CreatedAt,
		UpdatedAt:     accountRecord.UpdatedAt,
	}
}
