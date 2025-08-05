package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppEnv struct {
	GinModeDev  string
	GinModeProd string
	PORT        string
	DBHost      string
	DBPort      string
	DBUserName  string
	DBPassword  string
	DBName      string
}

var env *AppEnv

func GetEnv() AppEnv {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if env == nil {
		env = &AppEnv{
			PORT:        os.Getenv("PORT"),
			GinModeDev:  os.Getenv("GIN_MODE_DEV"),
			GinModeProd: os.Getenv("GIN_MODE_PROD"),
			DBHost:      os.Getenv("DB_HOST"),
			DBPort:      os.Getenv("DB_PORT"),
			DBUserName:  os.Getenv("DB_USERNAME"),
			DBPassword:  os.Getenv("DB_PASSWORD"),
			DBName:      os.Getenv("DB_NAME"),
		}
	}

	return *env
}
