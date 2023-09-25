package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Happy Path
func TestDataError(t *testing.T) {
	// JSON test
	jsonString := `{
        "id": 1,
        "messaje": "Error de datos"
    }`

	var dataErr DataError
	err := json.Unmarshal([]byte(jsonString), &dataErr)
	assert.NoError(t, err)


	assert.Equal(t, 1, dataErr.ID)
	assert.Equal(t, "Error de datos", dataErr.Message)

}

// Boundary cases
func TestWrongDataError(t *testing.T) {
	// JSON test
	jsonString := `{
        "id": "1",
        "messaje": "Error de datos"
    }`

	var dataErr DataError
	err := json.Unmarshal([]byte(jsonString), &dataErr)
	assert.Error(t, err)

}
