package project

import "part3/models/project/response"

type ProMod interface {
	ToProResponse() response.ProResponse
}
