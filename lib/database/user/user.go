package user

import (
	proResp "part3/models/project/response"
	taskResp "part3/models/task/response"
	"part3/models/user"
	"part3/models/user/request"
	"part3/models/user/response"

	"gorm.io/gorm"
)

type UserDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserDb {
	return &UserDb{db: db}
}

func (ud *UserDb) Create(newUser user.User) (user.User, error) {
	if err := ud.db.Create(&newUser).Error; err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (ud *UserDb) GetById(id int) (response.UserResponse, error) {
	userResp := response.UserResponse{}

	if err := ud.db.Model(&user.User{}).Where("id = ?", id).First(&userResp).Error; err != nil {
		return response.UserResponse{}, err
	}

	task := []taskResp.TaskResponse{}

	resTask := ud.db.Model(&user.User{}).Where("users.id = ?", id).Select("tasks.id as ID, tasks.created_at as CreatedAt, tasks.updated_at as UpdatedAt, tasks.name as Name, tasks.project_id as Project_id,tasks.priority as Priority ,projects.name as Project_name").Joins("inner join tasks on users.id = tasks.user_id").Joins("inner join projects on projects.id = tasks.project_id").Find(&task)

	if resTask.Error != nil {
		return userResp, resTask.Error
	}

	userResp.Tasks = task

	project := []proResp.ProResponse{}

	resPro := ud.db.Model(&user.User{}).Where("users.id = ?", id).Select("projects.id as Id, projects.created_at as Created_at, projects.updated_at as Updated_at, projects.name as Name").Joins("inner join projects on projects.user_id = users.id").Find(&project)

	if resPro.Error != nil {
		return userResp, resTask.Error
	}
	userResp.Projects = project

	return userResp, nil
}

func (ud *UserDb) UpdateById(id int, userReg request.UserRegister) (user.User, error) {

	_, err := ud.GetById(id)

	if err != nil {
		return user.User{}, err
	}

	ud.db.Model(&user.User{Model: gorm.Model{ID: uint(id)}}).Updates(user.User{Name: userReg.Name, Email: userReg.Email, Password: userReg.Password})

	user := userReg.ToUser()

	return user, nil
}

func (ud *UserDb) DeleteById(id int) (gorm.DeletedAt, error) {
	user := user.User{}
	_, err := ud.GetById(id)

	if err != nil {
		return user.DeletedAt, err
	}

	ud.db.Model(&user).Where("id = ?", id).Delete(&user)

	return user.DeletedAt, nil
}

func (ud *UserDb) GetAll() ([]response.UserResponse, error) {
	userRespArr := []response.UserResponse{}

	if err := ud.db.Model(user.User{}).Limit(5).Find(&userRespArr).Error; err != nil {
		return nil, err
	}

	for i := 0; i < len(userRespArr); i++ {
		task := []taskResp.TaskResponse{}
		resTask := ud.db.Model(&user.User{}).Where("users.id = ?", userRespArr[i].ID).Select("tasks.id as ID, tasks.created_at as CreatedAt, tasks.updated_at as UpdatedAt, tasks.name as Name, tasks.project_id as Project_id,tasks.priority as Priority ,projects.name as Project_name").Joins("inner join tasks on users.id = tasks.user_id").Joins("inner join projects on projects.id = tasks.project_id").Find(&task)

		if resTask.Error != nil {
			return userRespArr, resTask.Error
		}

		userRespArr[i].Tasks = task

		project := []proResp.ProResponse{}

		resPro := ud.db.Model(&user.User{}).Where("users.id = ?", userRespArr[i].ID).Select("projects.id as Id, projects.created_at as Created_at, projects.updated_at as Updated_at, projects.name as Name").Joins("inner join projects on projects.user_id = users.id").Find(&project)

		if resPro.Error != nil {
			return userRespArr, resTask.Error
		}
		userRespArr[i].Projects = project
	}

	return userRespArr, nil
}
