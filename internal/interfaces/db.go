package interfaces

import "github.com/Akanibekuly/tsarka-tz/internal/domain/entities"

type Db interface {
	Connect(dsn string) error
	Close()
	UserGet(id int) (*entities.UserSt, error)
	UserCreate(user *entities.UserSt) (int, error)
	UserUpdate(id int, user *entities.UserUpdateSt) error
	UserDelete(id int) error
}
