package router

import (

	"github.com/gin-gonic/gin"
	"net/http"

	"modak_ratelimit/api/handlers"
	"modak_ratelimit/api/middleware"
	"modak_ratelimit/internal/app/utils/logger"
)

// MapURL default url mapper.
func MapURL(router *gin.Engine) {

	logger.Info("Start router service")


	// Add health check
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running well")
	})


	// Supports version increment for endpoint
	v1 := router.Group("/v1")
	{
		v1.GET("/ratelimit", middleware.CheckRateLimit(), handlers.CheckRateLimit)

	}

}

