package user

import (
	"part3/models/user"
	"part3/models/user/request"

	"gorm.io/gorm"
)

type User interface {
	Create(newUser user.User) (user.User, error)
	GetById(id int) ([]user.User, error)
	UpdateById(id int, userid int, upUser request.UserRegister) (user.User, error)
	DeleteById(id int) (gorm.DeletedAt, error)
	GetAll() (user.User, error)
}
