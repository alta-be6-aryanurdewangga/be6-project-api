package project

import "part3/models/project"

type Project interface {
	Create(user_id int, newPro project.Project) (project.Project, error)
	GetById(id int, user_id int) (project.Project, error)
}
