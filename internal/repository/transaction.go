package repository

import (
	"context"
	"time"

	"github.com/go-rel/rel"
	"github.com/shopspring/decimal"
	"imansohibul.my.id/account-domain-service/entity"
)

type transactionRepository struct {
	db rel.Repository
}

type transaction struct {
	ID             uint            `db:"id"`
	AccountID      uint            `db:"account_id"`
	Type           int             `db:"type"`
	Amount         decimal.Decimal `db:"amount"`
	InitialBalance decimal.Decimal `db:"initial_balance"`
	FinalBalance   decimal.Decimal `db:"final_balance"`
	Currency       int             `db:"currency"`
	CreatedAt      time.Time       `db:"created_at"`
	UpdatedAt      time.Time       `db:"updated_at"`
}

func NewTransactionRepository(db rel.Repository) *transactionRepository {
	return &transactionRepository{db: db}
}

func (t transactionRepository) CreateTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	transactionRecord := t.fromEntityTransaction(transaction)
	err := t.db.Insert(ctx, transactionRecord)
	if err != nil {
		return nil, err
	}

	return t.toEntityTransaction(transactionRecord), nil
}

func (t transactionRepository) fromEntityTransaction(transactionEntity *entity.Transaction) *transaction {
	return &transaction{
		ID:             transactionEntity.ID,
		AccountID:      transactionEntity.AccountID,
		Type:           int(transactionEntity.Type),
		Amount:         transactionEntity.Amount,
		InitialBalance: transactionEntity.InitialBalance,
		FinalBalance:   transactionEntity.FinalBalance,
		Currency:       int(transactionEntity.Currency),
	}
}

func (t transactionRepository) toEntityTransaction(transactionRecord *transaction) *entity.Transaction {
	return &entity.Transaction{
		ID:             transactionRecord.ID,
		AccountID:      transactionRecord.AccountID,
		Type:           entity.TransactionType(transactionRecord.Type),
		Amount:         transactionRecord.Amount,
		InitialBalance: transactionRecord.InitialBalance,
		FinalBalance:   transactionRecord.FinalBalance,
		Currency:       entity.Currency(transactionRecord.Currency),
	}
}
