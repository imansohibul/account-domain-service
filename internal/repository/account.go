package repository

import "github.com/go-rel/rel"

type accountRepository struct {
	db rel.Repository
}

func NewAccountRepository(db rel.Repository) *accountRepository {
	return &accountRepository{db: db}
}
