package router

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"modak_ratelimit/api/handlers"
	"modak_ratelimit/api/middleware"

)


// MapURL default url mapper.
func MapURL(router *gin.Engine) {

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

