package entity


import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/go-playground/validator/v10"
)

// Happy Path
func TestInputData(t *testing.T) {
	// JSON test
	jsonString := `{
        "flow_id": "exampleflow",
        "user_id": "user123",
        "type": "alpha123",
        "lan": "en"
    }`


	var inputData InputData
	err := json.Unmarshal([]byte(jsonString), &inputData)
	assert.NoError(t, err)

	assert.Equal(t, "exampleflow", inputData.FlowID)
	assert.Equal(t, "user123", inputData.UserID)
	assert.Equal(t, "alpha123", inputData.Type)
	assert.Equal(t, "en", inputData.Lang)


	validate := validator.New()
	err = validate.Struct(inputData)
	assert.NoError(t, err)
}

// Boundary cases

func TestWrongInput(t *testing.T) {
	// JSON test
	jsonString := `{
        "flow_id": "exampleflow.#",
        "user_id": "user123",
        "type": "alpha123",
        "lan": "en"
    }`


	var inputData InputData
	err := json.Unmarshal([]byte(jsonString), &inputData)
	assert.NoError(t, err)

	validate := validator.New()
	err = validate.Struct(inputData)
	assert.Error(t, err)
}