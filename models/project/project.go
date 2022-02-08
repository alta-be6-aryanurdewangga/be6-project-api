package project

import "gorm.io/gorm"

type Project struct {
	gorm.Model

	Name_Pro string `json"name_pro"`
}
