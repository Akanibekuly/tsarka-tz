package interfaces

import "github.com/Akanibekuly/tsarka-tz/internal/domain/entities"

type Db interface {
	UserGet(id int) (*entities.UserSt, error)
	UserCreate(user *entities.UserSt) (int, error)
	UserUpdate(id int, user *entities.UserUpdateSt) error
	UserDelete(id int) error

	HashCreate(id string) error
	HashGet(id string) (*entities.HashSt, error)
	HashUpdate(id string, hash *entities.HashSt) error
}
