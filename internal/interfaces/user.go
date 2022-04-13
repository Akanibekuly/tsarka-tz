package interfaces

import "github.com/Akanibekuly/tsarka-tz/internal/domain/entities"

type User interface {
	Get(id int) (*entities.UserSt, error)
	Create(user *entities.UserSt) (int, error)
	Update(id int, user *entities.UserUpdateSt) error
	Delete(id int) error
}
