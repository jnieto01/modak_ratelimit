package server

import (

	"github.com/gin-gonic/gin"

	"modak_ratelimit/api/router"
	"modak_ratelimit/config"
	"modak_ratelimit/internal/app/utils/logger"
)

// Router exposes the endpoint of the application.

var Router *gin.Engine

// StartApp setup and run the app
func StartApp() {

	err := config.LoadConfig()
	if err != nil {
		logger.Error("Error with configuration file", err)
		return
	}


	setupRouter()
	runServer()
}

func setupRouter() {

	if config.App.Server.GoEnv == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}


	Router = gin.Default()
	router.MapURL(Router)
}

func runServer() {

	port := config.App.Server.Port
	host := config.App.Server.Host

	if err := Router.Run(host + ":" + port); err != nil {
		logger.Error("Error running server: ", err)
	}

}

