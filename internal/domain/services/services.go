package services

import (
	"github.com/Akanibekuly/tsarka-tz/internal/domain/entities"
	"github.com/Akanibekuly/tsarka-tz/internal/domain/repository"
	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
)

type Counter interface {
	Add(n int) (int, error)
	Sub(n int) (int, error)
	Val() (int, error)
	Del() error
}

type Hash interface {
	Calc(s []byte) (string, error)
	Result(id string) (*entities.HashSt, error)
}

type User interface {
	Get(id int) (*entities.UserSt, error)
	Create(user *entities.UserSt) (int, error)
	Update(id int, user *entities.UserUpdateSt) error
	Delete(id int) error
}

type Services struct {
	User
	Hash
	Counter
}

func New(lg interfaces.Logger, cache interfaces.Cache, reps *repository.Repository) *Services {
	return &Services{
		User: NewUserService(reps.User),
		Hash: NewHashService(lg, reps.Hash),
	}
}
