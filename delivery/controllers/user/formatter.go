package user

import "part3/models/user"

type GetUserResponseFormat struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    user.User `json:"data"`
}
