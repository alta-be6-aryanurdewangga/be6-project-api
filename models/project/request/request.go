package request

import "part3/models/project"

type ProRequest struct {
	Name string `json:"name"`
}

func (p *ProRequest) ToProject() project.Project {
	return project.Project{
		Name: p.Name,
	}
}
