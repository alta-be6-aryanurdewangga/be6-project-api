package project

import (
	"part3/configs"
	"part3/models/project"
	"part3/models/project/request"
	"part3/models/task"
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

	repo := New(db)

	t.Run("success run Create", func(t *testing.T) {
		mockPro := project.Project{User_ID: 2, Name_Pro: "anonim"}
		res, err := repo.Create(int(mockPro.User_ID), mockPro)
		log.Info(res.User_ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
		assert.Equal(t, 2, int(res.User_ID))
	})

	t.Run("fail run Create", func(t *testing.T) {
		mockPro := project.Project{Model: gorm.Model{ID: 1}, User_ID: 1, Name_Pro: "anonim"}
		_, err := repo.Create(int(mockPro.User_ID), mockPro)
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	confg := configs.GetConfig()
	db := utils.InitDB(confg)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	repo := New(db)
	mockCreate := project.Project{User_ID: 1, Name_Pro: "anonim"}
	_, err := repo.Create(int(mockCreate.User_ID), mockCreate)
	if err != nil {
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
	repo := New(db)
	mockCreate := project.Project{User_ID: 1, Name_Pro: "anonim"}
	_, err := repo.Create(int(mockCreate.User_ID), mockCreate)
	if err != nil {
		t.Fatal()
	}

	t.Run("success run UpdateById", func(t *testing.T) {
		mockPro := request.ProRequest{Name_Pro: "anonim321"}
		res, err := repo.UpdateById(1, 1, mockPro)
		assert.Nil(t, err)
		assert.Equal(t, "anonim321", res.Name_Pro)
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		mockPro := request.ProRequest{Name_Pro: "anonim321"}
		_, err := repo.UpdateById(2, 1, mockPro)
		assert.NotNil(t, err)
	})
}

func TestDeleteById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&project.Project{})
	repo := New(db)
	mockCreate := project.Project{User_ID: 1, Name_Pro: "anonim"}
	_, err := repo.Create(int(mockCreate.User_ID), mockCreate)
	if err != nil {
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
	repo := New(db)
	mockCreate := project.Project{User_ID: 1, Name_Pro: "anonim"}
	_, err := repo.Create(int(mockCreate.User_ID), mockCreate)
	if err != nil {
		t.Fatal()
	}

	t.Run("success run GetAll", func(t *testing.T) {
		res, err := repo.GetAll(1)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
	db.Migrator().DropTable(&project.Project{})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
	t.Run("fail run GetAll", func(t *testing.T) {
		_, err := repo.GetAll(1)
		assert.NotNil(t, err)
	})

}
