package handlers

import (

	"github.com/gin-gonic/gin"
	"net/http"

	"modak_ratelimit/internal/app/entity"
	"modak_ratelimit/internal/app/i18n"
	"modak_ratelimit/internal/app/usecase"
	"modak_ratelimit/internal/app/utils/logger"
)

func CheckRateLimit(c *gin.Context) {

	response := entity.Response{
		Status: http.StatusUnauthorized,
		Data: entity.ResponseData{
			IsAllowed: false,
			Error:     i18n.UnsupportedFlow,
		},
	}

	data := entity.InputData{
		FlowID: c.Query("flow_id"),
		UserID: c.Query("user_id"),
		Type:   c.Query("type"),
	}

	logger.Info("Start handler for check rate limited, user_id: " + data.UserID)

	switch data.FlowID {
	case "notifications":
		usecase.NotificacionsService(c, data)
		return
		/********************
		 Add any usecase like ratelit for vulnerability protection for all microservices
		*********************/
	}

	c.JSON(response.Status, response.Data)

}

