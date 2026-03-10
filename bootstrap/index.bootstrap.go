package bootstrap

import (
	"restapi-gin/config/app_config"
	"restapi-gin/database"
	"restapi-gin/routes"

	"github.com/gin-gonic/gin"
)

func BootstrapApp() {
	database.ConnectDatabase()
	app := gin.Default()

	routes.InitRoute(app)

	app.Run(app_config.PORT)
}
