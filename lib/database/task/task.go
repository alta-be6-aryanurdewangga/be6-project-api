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

func (td *TaskDb) Create(newTask task.Task) (task.Task, error) {
	if err := td.db.Create(&newTask).Error; err != nil {
		return newTask, err
	}
	return newTask, nil
}

func (td *TaskDb) GetById(id int) (task.Task, error) {
	task := task.Task{}

	if err := td.db.Model(&task).Where("id = ?", id).First(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (td *TaskDb) UpdateById(id int, taskReg request.TaskRequest) (task.Task, error) {
	_, err := td.GetById(id)

	if err != nil {
		return task.Task{}, err
	}

	td.db.Model(task.Task{Model: gorm.Model{ID: uint(id)}}).Updates(task.Task{Name_Task: taskReg.Name_Task, Priority: taskReg.Priority })

	task := taskReg.ToTask()

	return task, nil
}

func (bd *TaskDb) DeleteById(id int) (gorm.DeletedAt, error)  {
	task := task.Task{}
	_, err := bd.GetById(id)

	if err != nil{
		return task.DeletedAt, err
	}

	bd.db.Model(&task).Where("id = ?", id).Delete(&task)

	return task.DeletedAt, nil
}

func (bd *TaskDb) GetAll() ([]response.TaskResponse, error) {
	taskRespArr := []response.TaskResponse{}

	if err := bd.db.Model(task.Task{}).Limit(5).Find(&taskRespArr).Error ; err != nil {
		return nil, err
	}

	return taskRespArr, nil
}
