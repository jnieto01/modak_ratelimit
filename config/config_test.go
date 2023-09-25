package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
    "encoding/json"

    "modak_ratelimit/internal/app/entity"
)

//********** Unit Test ***********************

// Happy Path
func TestLoadConfig(t *testing.T){
	err:= LoadConfig()
    assert.Nil(t, err)
}

func TestValidationJson(t *testing.T){
    var AppTest entity.Config
    err := json.Unmarshal([]byte(jsonString), &AppTest)

    assert.Nil(t, err)	
}

func TestConfigVar(t *testing.T){
    err:= LoadConfig()
    assert.Nil(t, err)

	assert.Equal(t, App.Server.Port, "8080")
}

