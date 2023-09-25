package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Happy Path
func TestRedisDB(t *testing.T) {
	// JSON test
	jsonString := `{
        "addr": "localhost:6379",
        "password": "mypassword",
        "db": 1
    }`

	
	var redisDB RedisDB
	err := json.Unmarshal([]byte(jsonString), &redisDB)
	assert.NoError(t, err)

	assert.Equal(t, "localhost:6379", redisDB.Addr)
	assert.Equal(t, "mypassword", redisDB.Password)
	assert.Equal(t, 1, redisDB.Db)

}

// Boundary cases
func TestWrongDataRedisDB(t *testing.T) {
	// JSON test
	jsonString := `{
        "addr": 123,
        "password": "mypassword",
        "db": 1
    }`

	
	var redisDB RedisDB
	err := json.Unmarshal([]byte(jsonString), &redisDB)
	
	assert.Error(t, err)

}