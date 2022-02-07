package user

import (
	"part3/models/user/response"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt" `
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	Name     string `json:"name" gorm:"not null;type:varchar(100)"`
	Email    string `json:"email" gorm:"index;not null;type:varchar(100)"`
	Password string `json:"password" gorm:"not null;type:varchar(100)"`

}

func (u *User) ToUserResponse() response.UserResponse {
	return response.UserResponse{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		
		Name:      u.Name,
		Email:     u.Email,
	}
}
