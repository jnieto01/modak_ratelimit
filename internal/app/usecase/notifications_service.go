package usecase

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"modak_ratelimit/config"
	"modak_ratelimit/internal/app/entity"
	"modak_ratelimit/internal/app/i18n"
	"modak_ratelimit/internal/app/utils/logger"
)

func NotificacionsService(c *gin.Context, input entity.InputData) {

	logger.Info("Start notificacion service, user_id: " + input.UserID)

	response := entity.Response{
		Status: http.StatusOK,
		Data: entity.ResponseData{
			IsAllowed: false,
			Error:     i18n.SuspendedService,
		},
	}

	var rules []entity.RuleByType
	for _, flowRule := range config.App.RateLimitRules {
		if flowRule.FlowID == input.FlowID {
			rules = append(rules, flowRule.Settings...)
		}
	}

	var rule entity.RuleByType
	for _, inputRule := range rules {
		if inputRule.Key == input.Type {
			rule.Key = input.FlowID + "_" + inputRule.Key + "_" + input.UserID
			rule.MaxRequests = inputRule.MaxRequests
			rule.TimeInterval = inputRule.TimeInterval
		}
	}

	RateLimitExe(rule, &response)

	c.JSON(response.Status, response.Data)
}
