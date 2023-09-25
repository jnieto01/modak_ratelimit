package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Happy Path
func TestConfig(t *testing.T) {
	// JSON test
	jsonString := `{
        "Server": {
            "Host": "example.com",
            "Port": "8080",
            "GoEnv": "test"
        },
        "Kvs": {
            "Addr": "localhost:6379",
            "Password": "",
            "Db": 0
        },
        "RateLimitRules": [
            {
                "Flowid": "notifications",
                "Settings": [
                    {
                        "Key": "status",
                        "MaxRequests": 2,
                        "TimeInterval": 1
                    }
                ]
            }
        ]
    }`

	var cfg Config
	err := json.Unmarshal([]byte(jsonString), &cfg)
	assert.NoError(t, err)

	assert.Equal(t, "example.com", cfg.Server.Host)
	assert.Equal(t, "8080", cfg.Server.Port)
	assert.Equal(t, "test", cfg.Server.GoEnv)
	assert.Equal(t, "localhost:6379", cfg.Kvs.Addr)
	assert.Equal(t, "", cfg.Kvs.Password)
	assert.Equal(t, 0, cfg.Kvs.Db)

}

// Boundary cases