package repository

import (
	"context"
	"database/sql"
	"github.com/ilholbea/gym-rat/types"
)

type AccountRepository interface {
	Create(ctx context.Context, account types.Account) error
	GetByID(ctx context.Context, id int) (types.Account, error)
	Update(ctx context.Context, account types.Account) error
	Delete(ctx context.Context, id int) error
}

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(ctx context.Context, account types.Account) error {
	//TODO implement me
	panic("implement me")
}

func (r *accountRepository) GetByID(ctx context.Context, id int) (types.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (r *accountRepository) Update(ctx context.Context, account types.Account) error {
	//TODO implement me
	panic("implement me")
}

func (r *accountRepository) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
