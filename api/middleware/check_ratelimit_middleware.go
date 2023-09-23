package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"

	"modak_ratelimit/internal/app/utils/logger"
	"modak_ratelimit/internal/app/entity"
	"modak_ratelimit/internal/app/i18n"

)



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

		v := validator.New()

		data:= entity.InputData{
			FlowID: c.Query("flow_id"),
			UserID: c.Query("user_id"),
			Type: c.Query("type"),
			Lang: c.Query("lang"),
		}

		// language internationalization 
		i18n.SetLanguage(data.Lang)

		// Validation of input data 
		if err := v.Struct(data); err != nil {
			logger.Error("Middleware error:", err)

			response := entity.Response{
				IsAllowed: false,
				Error: i18n.ErrorMiddlewareQueryParams,
			}
				
			c.JSON(http.StatusBadRequest, response)
			c.Abort() 
			return
		}

		c.Next()
	}
}
