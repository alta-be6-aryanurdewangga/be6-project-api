package auth

import "part3/models/user"

type LoginRespFormat struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    user.User `json:"data"`
	Token   string    `json:"token"`
}
