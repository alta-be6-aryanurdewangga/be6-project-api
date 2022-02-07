package auth

import (
	"part3/models/user"
	"part3/models/user/request"
)

type Auth interface {
	Login(UserLogin request.Userlogin) (user.User, error)
}
