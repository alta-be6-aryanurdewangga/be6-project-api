package user

import (
	"part3/models/task"
	"part3/models/user/response"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string      `json:"name" gorm:"not null;type:varchar(100)"`
	Email    string      `json:"email" gorm:"unique;index;not null;type:varchar(100)"`
	Password string      `json:"password" gorm:"unique;not null;type:varchar(100)"`
	Tasks    []task.Task `json:"task" gorm:"foreignKey:User_ID"`
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
