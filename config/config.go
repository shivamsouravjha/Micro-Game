package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName      string
	SqlPrefix    string
	AppEnv       string
	DBUserName   string
	DBPassword   string
	DBHostWriter string
	DBHostReader string
	DBPort       string
	DBName       string
	JWT_SECRET   string
}

var config Config

func init() {
	appEnv := os.Getenv("APP_ENV")
	if len(appEnv) == 0 {
		appEnv = "dev"
	}

	configFilePath := "./config/.env"

	fmt.Println("reading env from: ", configFilePath)

	e := godotenv.Load(configFilePath)
	if e != nil {
		fmt.Println("error loading env: ", e)
		panic(e.Error())
	}
	config.SqlPrefix = "/* " + config.AppName + " - " + config.AppEnv + "*/"
	config.DBUserName = os.Getenv("DB_USERNAME")
	config.DBHostReader = os.Getenv("DB_HOST_READER")
	config.DBHostWriter = os.Getenv("DB_HOST_WRITER")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	config.JWT_SECRET = os.Getenv("JWT_SECRET")
}

func Get() Config {
	return config
}
