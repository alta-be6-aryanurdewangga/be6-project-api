package project

import (
	"part3/models/project/response"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model

	User_ID  uint   `json:"user_id"`
	Name_Pro string `json:"name_pro" gorm:"not null;type:varchar(100)"`
}

func (p *Project) ToProResponse() response.ProResponse {
	return response.ProResponse{
		Id:       p.ID,
		Name_Pro: p.Name_Pro,
	}
}
