package router

import (
 "testing"
 "github.com/stretchr/testify/assert"
 "net/http"
 "net/http/httptest"
 "github.com/gin-gonic/gin"
 "io/ioutil"
)

//********** Unit Test ***********************


// Happy Path
func TestServerRunning(t *testing.T){
	
	r := gin.Default()
	MapURL(r)
	
	req, _ := http.NewRequest("GET", "/ping", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error with the body: %v", err)
	}
		
	responseString := string(body)
	assert.Equal(t, "Server is running well", responseString, "Default valut must be: Server is running well" )	
}

//Boundary cases

func TestServerOffLine(t *testing.T){
	
	r := gin.Default()
		
	req, _ := http.NewRequest("GET", "/ping", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.NotEqual(t, http.StatusOK, resp.Code)

}
