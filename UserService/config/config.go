package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName              string
	DATABASE             string
	AppEnv               string
	SqlPrefix            string
	DBUserName           string
	PORT                 string
	DBPassword           string
	DBHostWriter         string
	DBHostReader         string
	DBPort               string
	DBName               string
	AMPQ_URL             string
	DBMaxOpenConnections int
	DBMaxIdleConnections int
}

var config Config

// Should run at the very beginning, before any other package
// or code.
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
	config.AppName = os.Getenv("SERVICE_NAME")
	config.AppEnv = appEnv
	config.SqlPrefix = "/* " + config.AppName + " - " + config.AppEnv + "*/"
	config.DBUserName = os.Getenv("DB_USERNAME")
	config.DBHostReader = os.Getenv("DB_HOST_READER")
	config.DBHostWriter = os.Getenv("DB_HOST_WRITER")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	config.DATABASE = os.Getenv("DATABASE")
	config.AMPQ_URL = os.Getenv("AMPQ_URL")
	config.PORT = os.Getenv("PORT")
	config.DBMaxIdleConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONENCTION"))
	config.DBMaxOpenConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
}

func Get() Config {
	return config
}
