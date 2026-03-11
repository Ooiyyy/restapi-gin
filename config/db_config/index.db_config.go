package db_config

import "os"

var DB_DRIVER = "mysql"
var DB_HOST = "127.0.0.1"
var DB_PORT = "3306"
var DB_NAME = "restapi_gin"

var DB_USER = "root"
var DB_PASSWORD = ""

func InitDatabaseConfig() {
	driverEnv := os.Getenv("DB_DRIVER")
	if driverEnv != "" {
		DB_DRIVER = driverEnv
	}

	hostEnv := os.Getenv("DB_HOST")
	if hostEnv != "" {
		DB_HOST = hostEnv
	}

	portEnv := os.Getenv("DB_PORT")
	if portEnv != "" {
		DB_PORT = portEnv
	}

	dbNameEnv := os.Getenv("DB_NAME")
	if dbNameEnv != "" {
		DB_NAME = dbNameEnv
	}

	dbUserEnv := os.Getenv("DB_USER")
	if dbUserEnv != "" {
		DB_USER = dbUserEnv
	}

	dbPasswordEnv := os.Getenv("DB_PASSWORD")
	if dbPasswordEnv != "" {
		DB_PASSWORD = dbPasswordEnv
	}
}
