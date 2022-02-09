package user

import "part3/models/user"

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email" `
	Password string `json:"password"`
}

type GetUserResponseFormat struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    user.User `json:"data"`
}
