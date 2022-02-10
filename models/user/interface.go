package user

import "part3/models/user/response"

type UserMod interface {
	ToUserResponse() response.UserResponse
}
