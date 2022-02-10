package project

import (
	"part3/models/project"
	"part3/models/project/request"
	"part3/models/project/response"

	"gorm.io/gorm"
)

type ProDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProDb {
	return &ProDb{db: db}
}

func (pd *ProDb) Create(user_id int, newPro project.Project) (project.Project, error) {
	newPro.User_ID = uint(user_id)
	if err := pd.db.Create(&newPro).Error; err != nil {
		return newPro, err
	}
	return newPro, nil
}

func (pd *ProDb) GetById(id int, user_id int) (project.Project, error) {
	pro := project.Project{}

	if err := pd.db.Model(&pro).Where("id = ? AND user_id = ?", id, user_id).First(&pro).Error; err != nil {
		return pro, err
	}
	return pro, nil
}

func (pd *ProDb) UpdateById(id int, user_id int, upPro request.ProRequest) (project.Project, error) {
	_, err := pd.GetById(id, user_id)

	if err != nil {
		return project.Project{}, err
	}

	pd.db.Model(project.Project{Model: gorm.Model{ID: uint(id)}, User_ID: uint(user_id)}).Updates(project.Project{Name: upPro.Name})

	pro := upPro.ToProject()

	return pro, nil
}

func (pd *ProDb) DeleteById(id int, user_id int) (gorm.DeletedAt, error) {
	pro := project.Project{}
	_, err := pd.GetById(id, user_id)
	if err != nil {
		return pro.DeletedAt, err
	}

	pd.db.Model(&pro).Where("id = ? AND user_id = ?", id, user_id).Delete(&pro)

	return pro.DeletedAt, err
}

func (pd *ProDb) GetAll(user_id int) ([]response.ProResponse, error) {
	proRespArr := []response.ProResponse{}

	if err := pd.db.Model(project.Project{}).Where("user_id = ?", user_id).Find(&proRespArr).Error; err != nil {
		return nil, err
	}

	return proRespArr, nil
}
