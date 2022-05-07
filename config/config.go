package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port             int
	ConnectionString string
	JWTSecret        string
}

func NewConfig() (Config, error) {
	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		return Config{}, err
	}
	return Config{
		Port:             port,
		ConnectionString: getEnv("CONNECTION_STRING", "root:123456@tcp(localhost:3306)/plantapi?charset=utf8&parseTime=True&loc=Local"),
		JWTSecret:        getEnv("JWT_SECRET", "RAHASIA"),
	}, err
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
