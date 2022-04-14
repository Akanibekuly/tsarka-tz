package core

import (
	"github.com/Akanibekuly/tsarka-tz/internal/domain/entities"
	"github.com/Akanibekuly/tsarka-tz/internal/interfaces"
)

type UserService struct {
	db interfaces.Db
}

func NewUserService(db interfaces.Db) *UserService {
	return &UserService{
		db: db,
	}
}

func (u *UserService) Get(id int) (*entities.UserSt, error) {
	return u.db.UserGet(id)
}

func (u *UserService) Create(user *entities.UserSt) (int, error) {
	return u.db.UserCreate(user)
}

func (u *UserService) Update(id int, user *entities.UserUpdateSt) error {
	return u.db.UserUpdate(id, user)
}

func (u *UserService) Delete(id int) error {
	return u.db.UserDelete(id)
}
