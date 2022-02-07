package auth

import (
	"part3/models/user"
	"part3/models/user/request"

	"gorm.io/gorm"
)

type AuthDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthDb {
	return &AuthDb{
		db: db,
	}
}

func (ad *AuthDb) Login(UserLogin request.Userlogin) (user.User, error) {
	user := user.User{}
	if err := ad.db.Model(&user).Where("email = ? AND password = ?", UserLogin.Email, UserLogin.Password).First(&user).Error ; err != nil {
		return user, err
	}

	return user, nil
}
