package auth

import (
	"part3/configs"
	_lib "part3/lib/database/user"
	"part3/models/user"
	"part3/models/user/request"
	"part3/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repo := New(db)
	db.Migrator().DropTable(&user.User{})
	db.AutoMigrate(&user.User{})

	t.Run("success run login", func(t *testing.T) {
		mockUser := user.User{Name: "anonim123", Email: "anonim@123", Password: "anonim123"}
		_, err := _lib.New(db).Create(mockUser)
		if err != nil {
			t.Fail()
		}
		mockLogin := request.Userlogin{Email: "anonim@123", Password: "anonim123"}
		res, err := repo.Login(mockLogin)
		assert.Nil(t, err)
		assert.Equal(t, "anonim@123", res.Email)
		assert.Equal(t, "anonim123", res.Password)
	})

	t.Run("fail run login", func(t *testing.T) {
		mockLogin := request.Userlogin{Email: "anonim@456", Password: "anonim456"}
		_, err := repo.Login(mockLogin)
		assert.NotNil(t, err)
	})

}
