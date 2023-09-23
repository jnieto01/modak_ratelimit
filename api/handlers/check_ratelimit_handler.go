package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type AnswerRateLimit struct {
	IsAllowed bool `json:"isallowed"`
	Error string `json:"error"`
}

func CheckRateLimit(c *gin.Context) {
	
	result := AnswerRateLimit{
		IsAllowed: true,
		Error: "",
	}
	
	c.JSON(http.StatusOK, result)

}







