package middlewares

import (
	"fmt"
	"part3/configs"
	"part3/models/user"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(u user.User) (string, error) {
	codes := jwt.MapClaims{
		"id":   u.ID,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"auth": true,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, codes)
	// fmt.Println(token)
	return token.SignedString([]byte(configs.JWT_SECRET))
}

func ExtractTokenId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
	if user.Valid {
		codes := user.Claims.(jwt.MapClaims)
		fmt.Println(codes)
		id := codes["id"].(float64)
		return id
	}
	return 0
}
