package middlewares

import (
	"errors"
	"log"
	"time"

	"github.com/GroupProject2-Kelompok4/BE/app/config"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(config.JWT),
		SigningMethod: "HS256",
	})
}

func CreateToken(id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(15 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(config.JWT))
	if err != nil {
		log.Println("generate jwt error ", err.Error())
		return "", nil
	}
	return result, err
}

func ExtractToken(e echo.Context) (string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		return id, nil
	}
	return "", errors.New("failed to extract jwt-token")
}
