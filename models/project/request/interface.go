package request

import "part3/models/project"

type ProReq interface {
	ToProject() project.Project
}
