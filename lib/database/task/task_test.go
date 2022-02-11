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

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
	db.AutoMigrate(&user.User{})

	t.Run("success run Create", func(t *testing.T) {
		mocUserP := user.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}

		if _, err := _lib.New(db).Create(mocUserP); err != nil {
			t.Fatal()
		}
		mockTask := task.Task{Name: "anonim123", Priority: 1}
		res, err := repo.Create(1, mockTask)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.User_ID))
		assert.Equal(t, "anonim123", res.Name)
		assert.Equal(t, 1, int(res.Priority))
	})

	t.Run("fail run Create", func(t *testing.T) {
		mocUserP := user.User{Name: "anonim2", Email: "anonim@2", Password: "anonim2"}
		if _, err := _lib.New(db).Create(mocUserP); err != nil {
			t.Fatal()
		}
		mockTaskP := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
		if _, err := repo.Create(1, mockTaskP); err != nil {
			t.Fatal()
		}
		mockTask := task.Task{Model: gorm.Model{ID: 1}, User_ID: 1, Name: "anonim123", Priority: 1}
		_, err := repo.Create(int(mockTask.User_ID), mockTask)
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
	db.AutoMigrate(&user.User{})

	t.Run("success run GetById", func(t *testing.T) {
		mocUserP := user.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}

		if _, err := _lib.New(db).Create(mocUserP); err != nil {
			t.Fatal()
		}
		mockTaskP := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
		if _, err := repo.Create(1, mockTaskP); err != nil {
			t.Fatal()
		}
		res, err := repo.GetById(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
		assert.Equal(t, 1, int(res.User_ID))
	})

	t.Run("fail run GetById", func(t *testing.T) {
		mocUserP := user.User{Name: "anonim2", Email: "anonim@2", Password: "anonim2"}

		if _, err := _lib.New(db).Create(mocUserP); err != nil {
			t.Fatal()
		}
		mockTaskP := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
		if _, err := repo.Create(1, mockTaskP); err != nil {
			t.Fatal()
		}
		_, err := repo.GetById(10, 1)
		assert.NotNil(t, err)
	})
}

func TestUpdateById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
	db.AutoMigrate(&user.User{})

	t.Run("success run UpdateById", func(t *testing.T) {
		mocUserP := user.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
		if _, err := _lib.New(db).Create(mocUserP); err != nil {
			t.Fatal()
		}
		mockTaskP := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
		if _, err := repo.Create(1, mockTaskP); err != nil {
			t.Fatal()
		}
		mockTask := request.TaskRequest{Name: "anonim321", Priority: 2}
		res, err := repo.UpdateById(1, 1, mockTask)
		assert.Nil(t, err)
		assert.Equal(t, "anonim321", res.Name)
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		mocUserP := user.User{Name: "anonim2", Email: "anonim@2", Password: "anonim2"}
		if _, err := _lib.New(db).Create(mocUserP); err != nil {
			t.Fatal()
		}
		mockTaskP := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
		if _, err := repo.Create(1, mockTaskP); err != nil {
			t.Fatal()
		}
		mockTask := request.TaskRequest{Name: "anonim321", Priority: 2}
		_, err := repo.UpdateById(10, 1, mockTask)
		assert.NotNil(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
	db.AutoMigrate(&user.User{})

	t.Run("success run DeleteById", func(t *testing.T) {
		mocUserP := user.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
		if _, err := _lib.New(db).Create(mocUserP); err != nil {
			t.Fatal()
		}
		mockTaskP := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
		if _, err := repo.Create(1, mockTaskP); err != nil {
			t.Fatal()
		}
		res, err := repo.DeleteById(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, true, res.Valid)
	})

	t.Run("fail run DeleteById", func(t *testing.T) {
		mocUserP := user.User{Name: "anonim2", Email: "anonim@2", Password: "anonim2"}
		if _, err := _lib.New(db).Create(mocUserP); err != nil {
			t.Fatal()
		}
		mockTaskP := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
		if _, err := repo.Create(1, mockTaskP); err != nil {
			t.Fatal()
		}
		_, err := repo.DeleteById(10, 1)
		assert.NotNil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
	db.AutoMigrate(&user.User{})

	t.Run("success run GetAll", func(t *testing.T) {
		mocUserP := user.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
		if _, err := _lib.New(db).Create(mocUserP); err != nil {
			t.Fatal()
		}
		mockPro := project.Project{Name: "Proanonim"}
		if _, err := _libPro.New(db).Create(1, mockPro); err != nil {
			t.Fatal()
		}
		mockTaskP := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
		if _, err := repo.Create(1, mockTaskP); err != nil {
			t.Fatal()
		}
		_, err := repo.GetAll(1)
		assert.Nil(t, err)
	})

	t.Run("fail run GetAll", func(t *testing.T) {
		_, errT := repo.DeleteById(1, 1)
		if errT != nil {
			t.Fail()
		}
		_, err := repo.GetAll(1)
		assert.NotNil(t, err)
	})
}

func TestGetByIdResp(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	db.AutoMigrate(&task.Task{})
	db.AutoMigrate(&user.User{})

	t.Run("Success GetByIdResp", func(t *testing.T) {
		mocUserP := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}

		if _, err := _lib.New(db).Create(mocUserP); err != nil {
			t.Fatal()
		}

		mockProP := project.Project{Name: "Proanonim"}
		if _, err := _libPro.New(db).Create(1, mockProP); err != nil {
			t.Fatal()
		}

		mockTaskP := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
		if _, err := repo.Create(1, mockTaskP); err != nil {
			t.Fatal()
		}
		res, err := repo.GetByIdResp(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
		assert.Equal(t, 1, int(res.Project_id))
	})

	t.Run("fail run GetByIdResp", func(t *testing.T) {
		_, err := repo.GetByIdResp(10, 1)
		assert.NotNil(t, err)
	})
}
