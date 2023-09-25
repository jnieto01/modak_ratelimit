package middleware

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"

	"modak_ratelimit/internal/app/entity"
	"modak_ratelimit/internal/app/i18n"
	"modak_ratelimit/internal/app/utils/logger"
)

func CheckRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {

		data := entity.InputData{
			FlowID: c.Query("flow_id"),
			UserID: c.Query("user_id"),
			Type:   c.Query("type"),
			Lang:   c.Query("lang"),
		}

		logger.Info("Start middleware for check rate limited, user_id: " + data.UserID)

		// language internationalization
		i18n.SetLanguage(data.Lang)

		// **************************
		//	Validation of input data
		//****************************/
		v := validator.New()
		if err := v.Struct(data); err != nil {
			logger.Error("Middleware error:", err)

			response := entity.Response{
				Status: http.StatusBadRequest,
				Data: entity.ResponseData{
					IsAllowed: false,
					Error:     i18n.ErrorMiddlewareQueryParams,
				},
			}

			c.JSON(response.Status, response.Data)
			c.Abort()
			return
		}

		// **************************
		//	Validation of Authentication
		// 	just an example where can use a token validation
		//****************************/
		/*
			if !isAuthenticated {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(),})
				c.Abort()
				return
			}
		*/

		c.Next()
	}
}
