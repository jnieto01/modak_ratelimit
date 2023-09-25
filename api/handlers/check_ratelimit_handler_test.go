package handlers

import (

	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"

)


// ********** Unit Test ***********************
var basePath = "/test"

// Happy Path

func Skip_TestFlowSupported(t *testing.T) {

	r := gin.Default()
	r.GET(basePath, func(c *gin.Context) {
		CheckRateLimit(c)
	})

	queryParams := "?flow_id=notifications&user_id=jd123&type=news"
	req, _ := http.NewRequest("GET", basePath+queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

}

// Boundary cases
func Skip_TestFlowNotSupported(t *testing.T) {

	r := gin.Default()
	r.GET(basePath, func(c *gin.Context) {
		CheckRateLimit(c)
	})

	queryParams := "?flow_id=another_fllow&user_id=jd123&type=news"
	req, _ := http.NewRequest("GET", basePath+queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusUnauthorized, resp.Code)

}
