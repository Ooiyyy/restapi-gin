package config

import (
	"restapi-gin/config/app_config"
	"restapi-gin/config/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
}
