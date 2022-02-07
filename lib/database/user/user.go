package user

import (
	"part3/models/user"
	"part3/models/user/request"
	"part3/models/user/response"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserDb {
	return &UserDb{db: db}
}

func (ud *UserDb) Create(newUser user.User) (user.User, error) {
	if err := ud.db.Create(&newUser).Error; err != nil {
		log.Warn("error in found database", err)
		return newUser, err
	}
	return newUser, nil
}

func (ud *UserDb) GetById(id int) (user.User, error) {
	user := user.User{}

	if err := ud.db.Model(&user).Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ud *UserDb) UpdateById(id int, userReg request.UserRegister) (user.User, error) {

	_, err := ud.GetById(id)

	if err != nil {
		return user.User{}, err
	}

	ud.db.Model(&user.User{ID: uint(id)}).Updates(user.User{Name: userReg.Name, Email: userReg.Email, Password: userReg.Password})

	user := userReg.ToUser()
	
	return user, nil
}

func (ud *UserDb) DeleteById(id int) (gorm.DeletedAt, error) {
	user := user.User{}
	_, err := ud.GetById(id)

	if err != nil {
		return user.DeletedAt, err
	}

	ud.db.Model(&user).Where("id = ?", id).Delete(&user)

	return user.DeletedAt, nil
}

func (ud *UserDb) GetAll() ([]response.UserResponse, error) {
	userRespArr := []response.UserResponse{}

	if err := ud.db.Model(user.User{}).Limit(5).Find(&userRespArr).Error; err != nil {
		return nil, err
	}

	return userRespArr, nil
}
