package service

import (
	"github.com/ilholbea/gym-rat/types"
)

type Database interface {
	Create(exercise *types.Exercise) error
	//Get() error
	//Update() error
	//Delete() error
}
