package middleware

import (
	"errors"
	"fmt"
	"plant-api/business/user"
	"plant-api/config"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthorizeJWT(secret string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey: []byte(secret),
	})
}

// Get and parse JWT from header
func ParseJWT(c echo.Context) (user user.User, err error) {
	// Get token
	token := c.Request().Header.Get("Authorization")
	arrToken := strings.Split(token, " ")
	if len(arrToken) < 2 {
		err = errors.New("header authorization invalid value")
		return user, err
	}

	// Parse token
	tokenJwt, err := jwt.Parse(arrToken[1], func(token *jwt.Token) (interface{}, error) {
		cfg, _ := config.NewConfig()
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		return user, err
	}

	if !tokenJwt.Valid {
		return user, err
	}

	// Store the payload
	payload := tokenJwt.Claims.(jwt.MapClaims)
	user.ID = uint(payload["id"].(float64))
	user.Role = payload["role"].(string)
	return user, nil
}