package user

import (
	"part3/configs"
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

	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&user.User{})

	repo := New(db)
	t.Run("success run Create", func(t *testing.T) {
		mocUser := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
		res, err := repo.Create(mocUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("fail run Create", func(t *testing.T) {
		mocUser := user.User{Model: gorm.Model{ID: 1}, Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
		_, err := repo.Create(mocUser)
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run GetById", func(t *testing.T) {
		res, err := repo.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("fail run GetById", func(t *testing.T) {
		res, err := repo.GetById(2)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, int(res.ID))
	})
}

func TestUpdateById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run UpdateById", func(t *testing.T) {
		mockUser := request.UserRegister{Name: "anonim321", Email: "anonim@321", Password: "anonim321"}
		res, err := repo.UpdateById(1, mockUser)
		assert.Nil(t, err)
		assert.Equal(t, "anonim321", res.Name)
	})

	t.Run("fail run UpdateById", func(t *testing.T) {
		mockUser := request.UserRegister{Name: "anonim456", Email: "anonim@456", Password: "456"}
		res, err := repo.UpdateById(2, mockUser)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, int(res.ID))
	})
}

func TestDeleteById(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)

	t.Run("success run DeleteById", func(t *testing.T) {
		res, err := repo.DeleteById(1)
		assert.Nil(t, err)
		assert.Equal(t, true, res.Valid)
	})

	t.Run("fail run DeleteById", func(t *testing.T) {
		res, err := repo.DeleteById(1)
		assert.NotNil(t, err)
		assert.Equal(t, false, res.Valid)
	})
}

func TestGetAll(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	t.Run("success run GetAll", func(t *testing.T) {
		repo := New(db)
		res, err := repo.GetAll()
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
	t.Run("fail run GetAll", func(t *testing.T) {
		db.Migrator().DropTable(&user.User{})
		db.Migrator().DropTable(&task.Task{})
		repo := New(db)
		_, err := repo.GetAll()
		assert.NotNil(t, err)
	})
}
