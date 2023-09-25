package usecase

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	//"modak_ratelimit/config"
	"modak_ratelimit/internal/app/entity"
	//"modak_ratelimit/internal/app/utils/kvs"
)

// Happy Path
// This unit test requires redis docker to be running
/*
func Skip_TestRateLimit(t *testing.T) {

	response := &entity.Response{
		Status: http.StatusOK,
		Data:   entity.ResponseData{},
	}

	rule := entity.RuleByType{
		Key:          "test_key",
		MaxRequests:  2,
		TimeInterval: 1,
	}

	DB := entity.RedisDB{
		Addr:     config.App.Kvs.Addr,
		Password: config.App.Kvs.Password,
		Db:       config.App.Kvs.Db,
	}

	con, err := kvs.NewClient(DB)
	assert.Nil(t, err)

	err = con.Delete(rule.Key)
	assert.Nil(t, err)
	con.Close()

	for i := 0; i < rule.MaxRequests + 5; i++ {

		if i < rule.MaxRequests {
			RateLimitExe(rule, response)

			assert.Equal(t, http.StatusOK, response.Status)
			assert.True(t, response.Data.IsAllowed)

		} else {
			RateLimitExe(rule, response)

			assert.Equal(t, http.StatusOK, response.Status)
			assert.False(t, response.Data.IsAllowed)
		}
	}


	con, err = kvs.NewClient(DB)
	assert.Nil(t, err)
	defer con.Close()

	err = con.Delete(rule.Key)
	assert.Nil(t, err)

}
*/

func TestInternalError(t *testing.T) {
	response := &entity.Response{
		Status: http.StatusOK,
		Data:   entity.ResponseData{},
	}

	internalError(response)
	assert.Equal(t, http.StatusInternalServerError, response.Status)
	assert.False(t, response.Data.IsAllowed)
}

// Boundary cases
