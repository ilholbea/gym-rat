package services

import (
	"context"
	"github.com/ilholbea/gym-rat/repository"
	"github.com/ilholbea/gym-rat/types"
)

type AccountService interface {
	GetAll(ctx context.Context) ([]types.Account, error)
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

func (as *accountService) GetAll(ctx context.Context) ([]types.Account, error) {
	return as.repo.GetAll(ctx)
}

func (as *accountService) Create(ctx context.Context, account types.Account) error {
	//TODO implement me
	panic("implement me")
}

func (as *accountService) GetByID(ctx context.Context, id int) (types.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (as *accountService) Update(ctx context.Context, account types.Account) error {
	//TODO implement me
	panic("implement me")
}

func (as *accountService) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
