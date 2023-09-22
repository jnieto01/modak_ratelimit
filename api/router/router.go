package router

import (
	"net/http"
	"github.com/gin-gonic/gin"

)


// MapURL default url mapper.
func MapURL(router *gin.Engine) {

	// Add health check
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running well")
	})

	
}

