package usecase

import (
	
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"

	"modak_ratelimit/internal/app/entity"
	

)

// Happy Path

func Skip_TestNotificacionsService(t *testing.T) {

	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		
		data := entity.InputData{
			FlowID: c.Query("flow_id"),
			UserID: c.Query("user_id"),
			Type:   c.Query("type"),
		}

		NotificacionsService(c, data)

	})

	queryParams := "?flow_id=notifications&user_id=jd123&type=news"
	req, _ := http.NewRequest("GET", "/test" +queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

}


// Boundary cases
