package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"plant-api/api/common"
	"plant-api/business/user"
	"plant-api/config"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Generate JWT
func GenerateJWT(id int, role, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
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

// Grant user with super role
func GrantSuper(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := ParseJWT(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
		}

		if claims.Role != "super" {
			return c.JSON(http.StatusForbidden, common.ForbiddenResponse())
		}

		return next(c)
	}
}

// Grant user with admin role
func GrantAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := ParseJWT(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
		}

		if !common.ValidateByRole("admin", claims.Role) {
			return c.JSON(http.StatusForbidden, common.ForbiddenResponse())
		}

		return next(c)
	}
}

// Grant user with the same id or super role
func GrantByIDOrSuper(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := ParseJWT(c)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequestResponse())
		}

		id, _ := strconv.Atoi(c.Param("id"))
		if !common.ValidateById(id, claims.ID, claims.Role) {
			return c.JSON(http.StatusForbidden, common.ForbiddenResponse())
		}

		return next(c)
	}
}
