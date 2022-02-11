package user

import (
	"part3/configs"
	_libPro "part3/lib/database/project"
	_libTask "part3/lib/database/task"
	"part3/models/project"
	"part3/models/task"
	"part3/models/user"
	"part3/models/user/request"
	"part3/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&user.User{})

	t.Run("success run Create", func(t *testing.T) {
		mocUser := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
		res, err := repo.Create(mocUser)
		assert.Nil(t, err)
		assert.Equal(t, "anonim123", res.Name)
		assert.Equal(t, "anonim@123", res.Email)
		assert.Equal(t, "anonim123", res.Password)

	})

	t.Run("fail run Create", func(t *testing.T) {
		mocUserP := user.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}
		if _, err := repo.Create(mocUserP); err != nil {
			t.Fatal()
		}
		mocUser := user.User{Model: gorm.Model{ID: 1}, Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
		_, err := repo.Create(mocUser)
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
		mocUser := user.User{Name: "anonim123", Email: "anonim@1", Password: "anonim1"}

		if _, err := repo.Create(mocUser); err != nil {
			t.Fatal()
		}

		mockPro := project.Project{Name: "Proanonim"}
		if _, err := _libPro.New(db).Create(1, mockPro); err != nil {
			t.Fatal()
		}

		mockTask := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
		if _, err := _libTask.New(db).Create(1, mockTask); err != nil {
			t.Fatal()
		}
		res, err := repo.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))

	})

	t.Run("fail run GetById", func(t *testing.T) {
		mocUser := user.User{Name: "anonim123", Email: "anonim@2", Password: "anonim12"}

		if _, err := repo.Create(mocUser); err != nil {
			t.Fatal()
		}

		mockPro := project.Project{Name: "Proanonim"}
		if _, err := _libPro.New(db).Create(1, mockPro); err != nil {
			t.Fatal()
		}

		mockTask := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
		if _, err := _libTask.New(db).Create(1, mockTask); err != nil {
			t.Fatal()
		}
		res, err := repo.GetById(10)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, int(res.ID))
	})

	// t.Run("fail run GetById project", func(t *testing.T) {
	// 	mocUserP := user.User{Name: "anonim123", Email: "anonim@3", Password: "anonim3"}

	// 	if _, err := repo.Create(mocUserP); err != nil {
	// 		t.Fatal()
	// 	}

	// 	mockTaskP := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
	// 	if _, err := _libTask.New(db).Create(1, mockTaskP); err != nil {
	// 		t.Fatal()
	// 	}

	// 	db.Migrator().DropTable(&project.Project{})
	// 	_, err := repo.GetById(1)
	// 	assert.NotNil(t, err)
	// })

	// t.Run("fail run GetById task", func(t *testing.T) {
	// 	mocUserP := user.User{Name: "anonim4", Email: "anonim@4", Password: "anonim4"}

	// 	if _, err := repo.Create(mocUserP); err != nil {
	// 		t.Fatal()
	// 	}
	// 	mockTaskP := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
	// 	if _, err := _libTask.New(db).Create(1, mockTaskP); err != nil {
	// 		t.Fatal()
	// 	}
	// 	db.Migrator().DropTable(&task.Task{})
	// 	_, err := repo.GetById(1)
	// 	assert.NotNil(t, err)
	// })
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
		mocUser := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}

		if _, err := repo.Create(mocUser); err != nil {
			t.Fatal()
		}
		mockUser := request.UserRegister{Name: "anonim321", Email: "anonim@321", Password: "anonim321"}
		res, err := repo.UpdateById(1, mockUser)
		assert.Nil(t, err)
		assert.Equal(t, "anonim321", res.Name)
		assert.Equal(t, "anonim@321", res.Email)
		assert.Equal(t, "anonim321", res.Password)
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		mockUser := request.UserRegister{Name: "anonim456", Email: "anonim@456", Password: "456"}
		_, err := repo.UpdateById(10, mockUser)
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
		mocUser := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}

		if _, err := repo.Create(mocUser); err != nil {
			t.Fatal()
		}
		res, err := repo.DeleteById(1)
		assert.Nil(t, err)
		assert.Equal(t, true, res.Valid)
	})

	t.Run("fail run DeleteById", func(t *testing.T) {
		_, err := repo.DeleteById(10)
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
		mocUser := user.User{Name: "anonim1", Email: "anonim@1", Password: "anonim1"}

		if _, err := repo.Create(mocUser); err != nil {
			t.Fatal()
		}

		mockPro := project.Project{Name: "Proanonim"}
		if _, err := _libPro.New(db).Create(1, mockPro); err != nil {
			t.Fatal()
		}

		mockTask := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
		if _, err := _libTask.New(db).Create(1, mockTask); err != nil {
			t.Fatal()
		}

		res, err := repo.GetAll()
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("fail run GetAll", func(t *testing.T) {
		if _, err := repo.DeleteById(1); err != nil {
			t.Fatal()
		}
		_, err := repo.GetAll()
		assert.NotNil(t, err)
	})

	// t.Run("fail run GetById project", func(t *testing.T) {
	// 	mocUserA := user.User{Name: "anonim2", Email: "anonim@2", Password: "anonim2"}

	// 	if _, err := repo.Create(mocUserA); err != nil {
	// 		t.Fatal()
	// 	}

	// 	mockTaskA := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
	// 	if _, err := _libTask.New(db).Create(1, mockTaskA); err != nil {
	// 		t.Fatal()
	// 	}
	// 	db.Migrator().DropTable(&project.Project{})
	// 	_, err := repo.GetAll()
	// 	assert.NotNil(t, err)
	// })

	// t.Run("fail run GetAll task", func(t *testing.T) {
	// 	mocUserP := user.User{Name: "anonim3", Email: "anonim@3", Password: "anonim3"}

	// 	if _, err := repo.Create(mocUserP); err != nil {
	// 		t.Fatal()
	// 	}

	// 	mockTaskP := task.Task{Name: "Taskanonim123", Priority: 5, Project_id: 1}
	// 	if _, err := _libTask.New(db).Create(1, mockTaskP); err != nil {
	// 		t.Fatal()
	// 	}
	// 	db.Migrator().DropTable(&task.Task{})
	// 	_, err := repo.GetAll()
	// 	assert.NotNil(t, err)
	// })

}
