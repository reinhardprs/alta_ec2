package appConfig

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	HttpPort  int
	DbConfig  ConfigDB
	Storage   string // "db" or "mem"
	JwtSecret string
}

type ConfigDB struct {
	Host     string
	User     string
	Password string
	Port     string
	Name     string
}

func NewConfig() (Config, error) {
	var cfg Config
	httpPort, err := strconv.Atoi(getEnv("HTTP_PORT", "8000"))
	if err != nil {
		return cfg, err
	}

	dbConfig := ConfigDB{
		Host:     os.Getenv("DATABASE_HOST"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Port:     getEnv("DATABASE_PORT", "3306"),
		Name:     os.Getenv("DATABASE_NAME"),
	}

	return Config{
		HttpPort:  httpPort,
		DbConfig:  dbConfig,
		Storage:   getEnv("STORAGE", "db"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}, nil
}

func (db ConfigDB) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
