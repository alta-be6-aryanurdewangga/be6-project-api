package task

import (
	"part3/configs"
	_libPro "part3/lib/database/project"
	_lib "part3/lib/database/user"
	"part3/models/project"
	"part3/models/task"
	"part3/models/task/request"
	"part3/models/user"
	"part3/utils"
	"testing"

	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
	repo := New(db)

	mockUser := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
	_, err := _lib.New(db).Create(mockUser)
	if err != nil {
		t.Fail()
	}
	mockPro := project.Project{Name: "Proanonim"}
	if _, err := _libPro.New(db).Create(1, mockPro); err != nil {
		t.Fatal()
	}

	t.Run("success run Create", func(t *testing.T) {
		mockTask := task.Task{Name: "anonim123", Priority: 1}
		res, err := repo.Create(1, mockTask)
		assert.Nil(t, err)
		log.Info(mockTask)
		log.Info(res.User_ID)
		assert.Equal(t, 1, int(res.ID))
		assert.Equal(t, 1, int(res.User_ID))
		assert.Equal(t, "anonim123", res.Name_Task)
		assert.Equal(t, 1, int(res.Priority))
	})

	t.Run("fail run Create", func(t *testing.T) {
		mockTask := task.Task{Model: gorm.Model{ID: 1}, User_ID: 1, Name: "anonim123", Priority: 1}
		_, err := repo.Create(int(mockTask.User_ID), mockTask)
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
	repo := New(db)

	mockUser := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
	_, err := _lib.New(db).Create(mockUser)
	if err != nil {
		t.Fail()
	}
	mockPro := project.Project{Name: "Proanonim"}
	if _, err := _libPro.New(db).Create(1, mockPro); err != nil {
		t.Fatal()
	}

	mockTask := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
	if _, err := repo.Create(1, mockTask); err != nil {
		log.Info(err)
		t.Fatal()
	}

	t.Run("success run GetById", func(t *testing.T) {
		res, err := repo.GetById(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
		assert.Equal(t, 1, int(res.User_ID))
	})

	t.Run("fail run GetById", func(t *testing.T) {
		res, err := repo.GetById(2, 1)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, int(res.ID))
	})
}

func TestUpdateById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
	repo := New(db)

	mockUser := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
	_, err := _lib.New(db).Create(mockUser)
	if err != nil {
		t.Fail()
	}
	mockPro := project.Project{Name: "Proanonim"}
	if _, err := _libPro.New(db).Create(1, mockPro); err != nil {
		t.Fatal()
	}

	mockTask := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
	if _, err := repo.Create(1, mockTask); err != nil {
		log.Info(err)
		t.Fatal()
	}

	t.Run("success run UpdateById", func(t *testing.T) {
		mockTask := request.TaskRequest{Name: "anonim321", Priority: 2}
		res, err := repo.UpdateById(1, 1, mockTask)
		assert.Nil(t, err)
		assert.Equal(t, "anonim321", res.Name)
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		mockTask := request.TaskRequest{Name: "anonim321", Priority: 2}
		res, err := repo.UpdateById(2, 1, mockTask)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, int(res.ID))
	})
}

func TestDeleteById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
	repo := New(db)

	mockUser := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
	_, err := _lib.New(db).Create(mockUser)
	if err != nil {
		t.Fail()
	}
	mockPro := project.Project{Name: "Proanonim"}
	if _, err := _libPro.New(db).Create(1, mockPro); err != nil {
		t.Fatal()
	}

	mockTask := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
	if _, err := repo.Create(1, mockTask); err != nil {
		log.Info(err)
		t.Fatal()
	}

	t.Run("success run DeleteById", func(t *testing.T) {
		res, err := repo.DeleteById(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, true, res.Valid)
	})

	t.Run("fail run DeleteById", func(t *testing.T) {
		res, err := repo.DeleteById(1, 1)
		assert.NotNil(t, err)
		assert.Equal(t, false, res.Valid)
	})
}

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
	repo := New(db)

	mockUser := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
	_, err := _lib.New(db).Create(mockUser)
	if err != nil {
		t.Fail()
	}
	mockPro := project.Project{Name: "Proanonim"}
	if _, err := _libPro.New(db).Create(1, mockPro); err != nil {
		t.Fatal()
	}

	mockTask := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
	if _, err := repo.Create(1, mockTask); err != nil {
		log.Info(err)
		t.Fatal()
	}

	t.Run("success run GetAll", func(t *testing.T) {
		res, err := repo.GetAll(1)
		assert.Nil(t, err)
		assert.NotNil(t, res[0])
	})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	t.Run("fail run GetAll", func(t *testing.T) {

		_, err := repo.GetAll(1)
		assert.NotNil(t, err)
	})
}



func TestGetByIdResp(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
	repo := New(db)

	mockUser := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
	_, err := _lib.New(db).Create(mockUser)
	if err != nil {
		t.Fail()
	}
	mockPro := project.Project{Name: "Proanonim"}
	if _, err := _libPro.New(db).Create(1, mockPro); err != nil {
		t.Fatal()
	}

	mockTask := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
	if _, err := repo.Create(1, mockTask); err != nil {
		log.Info(err)
		t.Fatal()
	}

	t.Run("Success GetByIdResp", func(t *testing.T) {
		res, err := repo.GetByIdResp(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
		assert.Equal(t, 1, int(res.Project_id))
	})

	t.Run("fail run GetByIdResp", func(t *testing.T) {
		res, err := repo.GetByIdResp(2, 1)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, int(res.ID))
	})
}