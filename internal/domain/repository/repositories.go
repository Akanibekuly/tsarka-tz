package repository

import (
	"github.com/Akanibekuly/tsarka-tz/internal/domain/entities"
	pg2 "github.com/Akanibekuly/tsarka-tz/internal/domain/repository/pg"
	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
	"github.com/jackc/pgx/v4"
)

type Repository struct {
	Hash
	User
}

type Hash interface {
	HashCreate(id string) error
	HashGet(id string) (*entities.HashSt, error)
	HashUpdate(id string, hash *entities.HashSt) error
}

type User interface {
	UserGet(id int) (*entities.UserSt, error)
	UserCreate(user *entities.UserSt) (int, error)
	UserUpdate(id int, user *entities.UserUpdateSt) error
	UserDelete(id int) error
}

func New(lg interfaces.Logger, conn *pgx.Conn) *Repository {
	return &Repository{
		User: pg2.NewUserRepository(lg, conn),
		Hash: pg2.NewHashRepository(lg, conn),
	}
}
