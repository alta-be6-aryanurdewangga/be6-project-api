package user

import (
	"part3/models/task"
	"part3/models/user/response"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// User_ID  int         `gorm:"autoIncrement"`
	Name     string      `gorm:"not null;type:varchar(100)"`
	Email    string      `gorm:"unique;index;not null;type:varchar(100)"`
	Password string      `gorm:"unique;not null;type:varchar(100)"`
	Tasks    []task.Task `gorm:"foreignKey:user_id"`
}

func (u *User) ToUserResponse() response.UserResponse {
	return response.UserResponse{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,

		Name:  u.Name,
		Email: u.Email,
	}
}
