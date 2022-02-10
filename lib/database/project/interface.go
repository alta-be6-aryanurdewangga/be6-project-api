package project

import (
	"part3/models/project"
	"part3/models/project/request"
	"part3/models/project/response"

	"gorm.io/gorm"
)

type Project interface {
	Create(user_id int, newPro project.Project) (project.Project, error)
	GetById(id int, user_id int) (project.Project, error)
	UpdateById(id int, user_id int, upPro request.ProRequest) (project.Project, error)
	DeleteById(id int, user_id int) (gorm.DeletedAt, error)
	GetAll(user_id int) ([]response.ProResponse, error)
}
