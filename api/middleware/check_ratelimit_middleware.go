package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"modak_ratelimit/internal/app/utils/logger"

)

type InputData struct {
	FlowID string `json:"flow_id" validate:"required,max=50,alphanum"`
    UserID string `json:"user_id" validate:"required,max=50,alphanum"`
    Type  string `json:"type" validate:"required,max=20,alphanum"`
}


func CheckRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		/*
		// auth required
		// just an example where can use a token validation
		if !isAuthenticated {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(),})
			c.Abort() 
			return
		}
		*/


		data:= InputData{
			FlowID: "",
			UserID: "",
			Type: "",
		}

		v := validator.New()

		data.FlowID = c.Query("flow_id")
		data.UserID = c.Query("user_id")
		data.Type = c.Query("type")
				
		if err := v.Struct(data); err != nil {
			logger.Error("Middleware error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),})
			c.Abort() 
			return
		}

		c.Next()
	}
}
