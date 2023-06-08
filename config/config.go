package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB *DbConfig
}

type DbConfig struct {
	Connection string
	Host       string
	Port       string
	Name       string
	Username   string
	Password   string
	Charset    string
}

func GetConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	return &Config{
		DB: &DbConfig{
			Connection: "mysql",
			Host:       dbHost,
			Port:       dbPort,
			Name:       dbName,
			Username:   dbUsername,
			Password:   dbPassword,
			Charset:    "utf8",
		},
	}
}
