package request

import "part3/models/project"

type ProRequest struct {
	Name_Pro string `json:"name_pro"`
}

func (p *ProRequest) ToProject() project.Project {
	return project.Project{
		Name_Pro: p.Name_Pro,
	}
}
