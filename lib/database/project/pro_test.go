package project

import (
	"part3/configs"
	_lib "part3/lib/database/user"
	"part3/models/project"
	"part3/models/project/request"
	"part3/models/task"
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

	t.Run("success run Create", func(t *testing.T) {
		db.Migrator().DropTable(&project.Project{})
		db.Migrator().DropTable(&task.Task{})
		db.Migrator().DropTable(&user.User{})
		db.AutoMigrate(&project.Project{})
		db.AutoMigrate(&task.Task{})

		mockUser := user.User{Name: "Useranonim123", Email: "anonim@123", Password: "anonim123"}

		if _, err := _lib.New(db).Create(mockUser); err != nil {
			t.Fatal()
		}
		mockPro := project.Project{Name: "Proanonim"}
		res, err := repo.Create(1, mockPro)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
		assert.Equal(t, 1, int(res.User_ID))
	})

	t.Run("fail run Create", func(t *testing.T) {
		db.Migrator().DropTable(&project.Project{})
		db.Migrator().DropTable(&task.Task{})
		db.Migrator().DropTable(&user.User{})
		mockPro := project.Project{Model: gorm.Model{ID: 1}, User_ID: 1, Name: "anonim"}
		_, err := repo.Create(int(mockPro.User_ID), mockPro)
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	confg := configs.GetConfig()
	db := utils.InitDB(confg)
	repo := New(db)

	t.Run("success run GetById", func(t *testing.T) {
		db.Migrator().DropTable(&project.Project{})
		db.Migrator().DropTable(&task.Task{})
		db.Migrator().DropTable(&user.User{})
		db.AutoMigrate(&project.Project{})
		db.AutoMigrate(&task.Task{})

		mockUser := user.User{Name: "Useranonim123", Email: "anonim@123", Password: "anonim123"}

		if _, err := _lib.New(db).Create(mockUser); err != nil {
			t.Fatal()
		}

		mockCreate := project.Project{Name: "anonim"}
		_, err := repo.Create(1, mockCreate)
		if err != nil {
			t.Fatal()
		}

		res, err := repo.GetById(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
		assert.Equal(t, 1, int(res.User_ID))
	})

	t.Run("fail run GetById", func(t *testing.T) {
		db.Migrator().DropTable(&project.Project{})
		db.Migrator().DropTable(&task.Task{})
		db.Migrator().DropTable(&user.User{})
		_, err := repo.GetById(2, 1)
		assert.NotNil(t, err)
	})
}

func TestUpdateById(t *testing.T) {
	confg := configs.GetConfig()
	db := utils.InitDB(confg)
	repo := New(db)

	t.Run("success run UpdateById", func(t *testing.T) {
		db.Migrator().DropTable(&project.Project{})
		db.Migrator().DropTable(&task.Task{})
		db.Migrator().DropTable(&user.User{})
		db.AutoMigrate(&project.Project{})
		db.AutoMigrate(&task.Task{})

		mockUser := user.User{Name: "Useranonim123", Email: "anonim@123", Password: "anonim123"}

		if _, err := _lib.New(db).Create(mockUser); err != nil {
			t.Fatal()
		}

		mockCreate := project.Project{Name: "anonim"}
		_, err := repo.Create(1, mockCreate)
		if err != nil {
			t.Fatal()
		}
		mockPro := request.ProRequest{Name: "anonim321"}
		res, err := repo.UpdateById(1, 1, mockPro)
		assert.Nil(t, err)
		assert.Equal(t, "anonim321", res.Name)
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		db.Migrator().DropTable(&project.Project{})
		db.Migrator().DropTable(&task.Task{})
		db.Migrator().DropTable(&user.User{})
		mockPro := request.ProRequest{Name: "anonim321"}
		_, err := repo.UpdateById(2, 1, mockPro)
		assert.NotNil(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	confg := configs.GetConfig()
	db := utils.InitDB(confg)
	repo := New(db)

	t.Run("success run DeleteById", func(t *testing.T) {
		db.Migrator().DropTable(&project.Project{})
		db.Migrator().DropTable(&task.Task{})
		db.Migrator().DropTable(&user.User{})
		db.AutoMigrate(&project.Project{})
		db.AutoMigrate(&task.Task{})

		mockUser := user.User{Name: "Useranonim123", Email: "anonim@123", Password: "anonim123"}

		if _, err := _lib.New(db).Create(mockUser); err != nil {
			t.Fatal()
		}

		mockCreate := project.Project{Name: "anonim"}
		_, err := repo.Create(1, mockCreate)
		if err != nil {
			t.Fatal()
		}
		res, err := repo.DeleteById(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, true, res.Valid)
	})

	t.Run("fail run DeleteById", func(t *testing.T) {
		db.Migrator().DropTable(&project.Project{})
		db.Migrator().DropTable(&task.Task{})
		db.Migrator().DropTable(&user.User{})
		_, err := repo.DeleteById(1, 1)
		assert.NotNil(t, err)
	})
}

func TestGetAll(t *testing.T) {
	confg := configs.GetConfig()
	db := utils.InitDB(confg)
	repo := New(db)

	t.Run("success run GetAll", func(t *testing.T) {
		db.Migrator().DropTable(&project.Project{})
		db.Migrator().DropTable(&task.Task{})
		db.Migrator().DropTable(&user.User{})
		db.AutoMigrate(&project.Project{})
		db.AutoMigrate(&task.Task{})

		mockUser := user.User{Name: "Useranonim123", Email: "anonim@123", Password: "anonim123"}

		if _, err := _lib.New(db).Create(mockUser); err != nil {
			t.Fatal()
		}

		mockCreate := project.Project{Name: "anonim"}
		_, err := repo.Create(1, mockCreate)
		if err != nil {
			t.Fatal()
		}
		res, err := repo.GetAll(1)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
	t.Run("fail run GetAll", func(t *testing.T) {
		db.Migrator().DropTable(&project.Project{})
		db.Migrator().DropTable(&task.Task{})
		db.Migrator().DropTable(&user.User{})
		_, err := repo.GetAll(1)
		assert.NotNil(t, err)
	})

}
