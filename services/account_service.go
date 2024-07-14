package services

import (
	"context"
	"github.com/ilholbea/gym-rat/repository"
	"github.com/ilholbea/gym-rat/types"
)

type AccountService interface {
	Create(ctx context.Context, account types.Account) error
	GetByID(ctx context.Context, id int) (types.Account, error)
	Update(ctx context.Context, account types.Account) error
	Delete(ctx context.Context, id int) error
}

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{repo: repo}
}

func (a accountService) Create(ctx context.Context, account types.Account) error {
	//TODO implement me
	panic("implement me")
}

func (a accountService) GetByID(ctx context.Context, id int) (types.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (a accountService) Update(ctx context.Context, account types.Account) error {
	//TODO implement me
	panic("implement me")
}

func (a accountService) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
