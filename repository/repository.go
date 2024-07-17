package repository

import (
	"context"
	"database/sql"

	"github.com/ilholbea/gym-rat/types"
)

type AccountRepository interface {
	GetAll(ctx context.Context) ([]types.Account, error)
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

func (ar *accountRepository) GetAll(ctx context.Context) ([]types.Account, error) {
	var accounts []types.Account
	rows, err := ar.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var account types.Account
		if err := rows.Scan(&account.ID, &account.FullName, &account.Email, &account.Password, &account.Type); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (ar *accountRepository) Create(ctx context.Context, account types.Account) error {
	//TODO implement me
	panic("implement me")
}

func (ar *accountRepository) GetByID(ctx context.Context, id int) (types.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (ar *accountRepository) Update(ctx context.Context, account types.Account) error {
	//TODO implement me
	panic("implement me")
}

func (ar *accountRepository) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
