package project

import (
	"part3/models/project/response"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model

	User_ID  uint
	Name_Pro string `gorm:"not null;type:varchar(100)"`
}

func (p *Project) ToProResponse() response.ProResponse {
	return response.ProResponse{
		Id:         p.ID,
		Created_at: p.CreatedAt,
		Updated_at: p.UpdatedAt,
		Name_Pro:   p.Name_Pro,
	}
}
