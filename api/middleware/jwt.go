package middleware

import (
	"fmt"
	"net/http"
	"plant-api/api/common"
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

type Claims struct {
	jwt.StandardClaims
	ID   int    `json:"id"`
	Role string `json:"role"`
	Exp  int64  `json:"exp"`
}

// Get and parse JWT from header
func ParseJWT(c echo.Context) (*Claims, error) {
	// Get token
	header := c.Request().Header.Get("Authorization")
	token := strings.Split(header, " ")[1]

	// Parse token
	tokenJwt, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		cfg, _ := config.NewConfig()
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenJwt.Claims.(*Claims); ok && tokenJwt.Valid {
		return claims, nil
	}
	return nil, err
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

		if claims.Role != "super" && claims.Role != "admin" {
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
		if int(claims.ID) != id && claims.Role != "super" {
			return c.JSON(http.StatusForbidden, common.ForbiddenResponse())
		}

		return next(c)
	}
}
