package repository

import (
	"context"

	"github.com/go-rel/rel"
)

type transactionManager struct {
	db rel.Repository
}

func NewTransactionManager(db rel.Repository) *transactionManager {
	return &transactionManager{db: db}
}

func (t transactionManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.db.Transaction(ctx, func(ctx context.Context) error {
		return fn(ctx)
	})
}
