package middlewares

import (
	"errors"
	"part3/configs"
	"part3/models/user"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(u user.User) (string, error) {
	if u.ID == 0 {
		return "cannot Generate token", errors.New("id == 0")
	}

	codes := jwt.MapClaims{
		"id":       u.ID,
		"email":     u.Email,
		"password": u.Password,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
		"auth":     true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, codes)
	// fmt.Println(token)
	return token.SignedString([]byte(configs.JWT_SECRET))
}

func ExtractTokenId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
	if user.Valid {
		codes := user.Claims.(jwt.MapClaims)
		id := codes["id"].(float64)
		return id
	}
	return 0
}

func ExtractTokenAdmin(e echo.Context) (result [2]string) {
	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
	if user.Valid {
		codes := user.Claims.(jwt.MapClaims)
		result[0] = codes["email"].(string)
		result[1] = codes["password"].(string)
		return result
	}
	return [2]string{}
}
