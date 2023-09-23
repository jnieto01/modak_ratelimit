package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"modak_ratelimit/internal/app/entity"
	"modak_ratelimit/internal/app/i18n"
)


func CheckRateLimit(c *gin.Context) {
	
	response := entity.Response{
		IsAllowed: true,
		Error: i18n.NotError,
	}


	
	c.JSON(http.StatusOK, response)

}







