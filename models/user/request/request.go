package request

import "part3/models/user"

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email" `
	Password string `json:"password"`
}

func (u *UserRegister) ToUser() user.User {
	return user.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
