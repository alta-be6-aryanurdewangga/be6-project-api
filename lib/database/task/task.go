package task

import (
	"part3/models/task"
	"part3/models/task/request"
	"part3/models/task/response"

	"gorm.io/gorm"
)

type TaskDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TaskDb {
	return &TaskDb{db: db}
}

func (td *TaskDb) Create(user_id int, newTask task.Task) (task.Task, error) {
	newTask.User_ID = uint(user_id)
	if err := td.db.Create(&newTask).Error; err != nil {
		return newTask, err
	}

	return newTask, nil
}

func (td *TaskDb) GetById(id int, user_id int) (task.Task, error) {
	task := task.Task{}

	if err := td.db.Model(&task).Where("id = ? AND user_id = ?", id, user_id).First(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (td *TaskDb) UpdateById(id int, user_id int, taskReg request.TaskRequest) (task.Task, error) {
	_, err := td.GetById(id, user_id)

	if err != nil {
		return task.Task{}, err
	}

	td.db.Model(task.Task{Model: gorm.Model{ID: uint(id)}, User_ID: uint(user_id)}).Updates(task.Task{Name: taskReg.Name, Priority: taskReg.Priority, Project_id: taskReg.Project_id})

	task := taskReg.ToTask()

	return task, nil
}

func (bd *TaskDb) DeleteById(id int, user_id int) (gorm.DeletedAt, error) {
	task := task.Task{}
	_, err := bd.GetById(id, user_id)

	if err != nil {
		return task.DeletedAt, err
	}

	bd.db.Model(&task).Where("id = ? AND user_id = ?", id, user_id).Delete(&task)

	return task.DeletedAt, nil
}

func (bd *TaskDb) GetAll(user_id int) ([]response.TaskResponse, error) {
	taskRespArr := []response.TaskResponse{}

	res := bd.db.Model(task.Task{}).Select("tasks.id as ID, tasks.created_at as CreatedAt, tasks.updated_at as UpdatedAt, tasks.name as Name, tasks.project_id as Project_id,tasks.priority as Priority ,projects.name as Project_name").Joins("inner join projects on projects.id = tasks.project_id").Find(&taskRespArr)
	if res.Error != nil {
		return nil, res.Error
	}
	return taskRespArr, nil
}

func (td *TaskDb) GetByIdResp(id int, user_id int) (response.TaskResponse, error) {
	taskResp := response.TaskResponse{}

	res := td.db.Model(task.Task{}).Where("tasks.id = ? AND tasks.user_id = ?", id, user_id).Select("tasks.id as ID, tasks.created_at as CreatedAt, tasks.updated_at as UpdatedAt, tasks.name as Name, tasks.project_id as Project_id,tasks.priority as Priority ,projects.name as Project_name").Joins("inner join projects on projects.id = tasks.project_id").First(&taskResp)

	if res.Error != nil {
		return response.TaskResponse{}, res.Error
	}

	return taskResp, nil
}
