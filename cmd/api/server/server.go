package server

import (
	"os"
	"github.com/gin-gonic/gin"
	
	"modak_ratelimit/api/router"
	"modak_ratelimit/internal/app/utils/logger"
)

//Router exposes the endpoint of the application.
var Router *gin.Engine

// StartApp setup and run the app
func StartApp() {
	setupRouter()
	runServer()
}

func setupRouter() {
	if env:=os.Getenv("GO_ENV"); env=="PROD" {
		gin.SetMode(gin.ReleaseMode)
	}else{
		gin.SetMode(gin.DebugMode)  
	}
	
	Router = gin.Default()
	router.MapURL(Router)
}

func runServer() {
	
	port := os.Getenv("PORT")	
	if port == "" {
		port = "8080"
	}

	if err := Router.Run(":"+port); err != nil {
		logger.Error("Error running server: " , err)
	}

}