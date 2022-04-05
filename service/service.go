package service

import (
	"github.com/ilholbea/gym-rat/types"
)

type Database interface {
	Create(exercise *types.Exercise) error
	GetAll() ([]types.Exercise, error)
	Get(id string) (*types.Exercise, error)
	Update(exercise *types.Exercise) error
	Delete(id string) error
}
