package auth

import (
	"part3/configs"
	_lib "part3/lib/database/user"
	"part3/models/task"
	"part3/models/user"
	"part3/models/user/request"
	"part3/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&user.User{})
	repo := New(db)

	mockUser := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
	_, err := _lib.New(db).Create(mockUser)
	if err != nil {
		t.Fail()
	}

	t.Run("success run login", func(t *testing.T) {
		mockLogin := request.Userlogin{Email: "anonim@123", Password: "anonim123"}
		res, err := repo.Login(mockLogin)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("fail run login", func(t *testing.T) {
		mockLogin := request.Userlogin{Email: "anonim@456", Password: "anonim456"}
		res, err := repo.Login(mockLogin)
		assert.NotNil(t, err)
		assert.NotEqual(t, 1, int(res.ID))
	})
	db.Migrator().DropTable(&task.Task{})
	db.Migrator().DropTable(&user.User{})
}
