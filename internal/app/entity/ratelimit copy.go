package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {

	jsonString := `{
        "status": 200,
        "data": {
            "isallowed": true,
            "error": {
                "id": 1,
                "message": "Error de datos"
            }
        }
    }`


	var resp Response
	err := json.Unmarshal([]byte(jsonString), &resp)
	assert.NoError(t, err)


	assert.Equal(t, 200, resp.Status)
	assert.True(t, resp.Data.IsAllowed)
	assert.Equal(t, 1, resp.Data.Error.ID)
	assert.Equal(t, "Error de datos", resp.Data.Error.Message)

}

func TestRateLimitRule(t *testing.T) {

	jsonString := `{
        "flowid": "example_flow",
        "settings": [
            {
                "key": "status",
                "maxrequests": 2,
                "timeinterval": 1
            },
            {
                "key": "news",
                "maxrequests": 1,
                "timeinterval": 1440
            }
        ]
    }`


	var rule RateLimitRule
	err := json.Unmarshal([]byte(jsonString), &rule)
	assert.NoError(t, err)

	assert.Equal(t, "example_flow", rule.FlowID)
	assert.Len(t, rule.Settings, 2)

}

func TestRuleByType(t *testing.T) {

	jsonString := `{
        "key": "status",
        "maxrequests": 2,
        "timeinterval": 1
    }`


	var rule RuleByType
	err := json.Unmarshal([]byte(jsonString), &rule)
	assert.NoError(t, err)

	assert.Equal(t, "status", rule.Key)
	assert.Equal(t, 2, rule.MaxRequests)
	assert.Equal(t, 1, rule.TimeInterval)

}
