package middleware

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strconv"
)

// ********** Unit Test ***********************
var basePath = "/test"

// Happy Path
func TestGoodParams(t *testing.T) {

	r := gin.Default()
	r.GET(basePath, CheckRateLimit())

	queryParams := "?flow_id=notifications&user_id=jd123&type=news"
	req, _ := http.NewRequest("GET", basePath+queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

}

//Boundary cases

func TestEmptyFlowID(t *testing.T) {

	r := gin.Default()
	r.GET(basePath, CheckRateLimit())

	queryParams := "?&user_id=jd123&type=news"
	req, _ := http.NewRequest("GET", basePath+queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestWrongFlowID(t *testing.T) {

	r := gin.Default()
	r.GET(basePath, CheckRateLimit())

	wrongFlowID := "notifications#.123"
	queryParams := "?flow_id=" + wrongFlowID + "&user_id=jd123&type=news"
	req, _ := http.NewRequest("GET", basePath+queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestTooLongFlowID(t *testing.T) {

	r := gin.Default()
	r.GET(basePath, CheckRateLimit())

	longFlowID := ""
	for i := 0; i < 60; i++ {
		longFlowID = longFlowID + strconv.Itoa(i)
	}

	queryParams := "?flow_id=" + longFlowID + "&user_id=jd123&type=news"
	req, _ := http.NewRequest("GET", basePath+queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestEmptyUserID(t *testing.T) {

	r := gin.Default()
	r.GET(basePath, CheckRateLimit())

	queryParams := "?flow_id=notifications&type=news"
	req, _ := http.NewRequest("GET", basePath+queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestWrongUserID(t *testing.T) {

	r := gin.Default()
	r.GET(basePath, CheckRateLimit())

	wrongUserID := "notifications#.123"
	queryParams := "?flow_id=" + wrongUserID + "&user_id=jd123&type=news"
	req, _ := http.NewRequest("GET", basePath+queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestTooLongUserID(t *testing.T) {

	r := gin.Default()
	r.GET(basePath, CheckRateLimit())

	longUserID := ""
	for i := 0; i < 60; i++ {
		longUserID = longUserID + strconv.Itoa(i)
	}

	queryParams := "?flow_id=" + longUserID + "&user_id=jd123&type=news"
	req, _ := http.NewRequest("GET", basePath+queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestEmptyType(t *testing.T) {

	r := gin.Default()
	r.GET(basePath, CheckRateLimit())

	queryParams := "?flow_id=notifications&user_id=jd123"
	req, _ := http.NewRequest("GET", basePath+queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestWrongType(t *testing.T) {

	r := gin.Default()
	r.GET(basePath, CheckRateLimit())

	wrongType := "notifications#.123"
	queryParams := "?flow_id=" + wrongType + "&user_id=jd123&type=news"
	req, _ := http.NewRequest("GET", basePath+queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestTooLongType(t *testing.T) {

	r := gin.Default()
	r.GET(basePath, CheckRateLimit())

	longType := ""
	for i := 0; i < 60; i++ {
		longType = longType + strconv.Itoa(i)
	}

	queryParams := "?flow_id=" + longType + "&user_id=jd123&type=news"
	req, _ := http.NewRequest("GET", basePath+queryParams, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
