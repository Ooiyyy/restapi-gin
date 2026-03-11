package bootstrap

import (
	"log"
	"restapi-gin/config"
	"restapi-gin/config/app_config"
	"restapi-gin/database"
	"restapi-gin/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {
	//Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	//Init Config
	config.InitConfig()

	//db connection
	database.ConnectDatabase()
	// Init gin engine
	app := gin.Default()
	// Initroute
	routes.InitRoute(app)
	// Run App
	app.Run(app_config.PORT)
}
