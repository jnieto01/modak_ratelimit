package i18n

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//********** Unit Test ***********************

//Boundary cases

func TestSetLanguage(t *testing.T) {
	err := SetLanguage("en")
	assert.Nil(t, err)
}

func TestLoadWithoutInicialization(t *testing.T) {

	assert.Equal(t, "", NotError.Message, "Default valut must be: ")
	assert.Equal(t, "Error in the request parameters", ErrorMiddlewareQueryParams.Message, "Default valut must be: Error in the request parameters")
}

func TestLoadEmpty(t *testing.T) {
	err := SetLanguage("")

	assert.Nil(t, err)
	assert.Equal(t, "", NotError.Message, "Default valut must be: ")
	assert.Equal(t, "Error in the request parameters", ErrorMiddlewareQueryParams.Message, "Default valut must be: Error in the request parameters")
}

func TestLangFile(t *testing.T) {
	// language not supported
	// language file does not exist
	err := SetLanguage("ar")

	assert.Nil(t, err)
	assert.Equal(t, "", NotError.Message, "Default valut must be: ")
	assert.Equal(t, "Error in the request parameters", ErrorMiddlewareQueryParams.Message, "Default valut must be: Error in the request parameters")
}

func TestLangKeyMissing(t *testing.T) {
	// key missing
	err := SetLanguage("en")

	assert.Nil(t, err)
	assert.Equal(t, "MessageMissing key is not in the file of language", MessageMissing.Message, "MessageMissing key is not in the file of language")

}

// Happy Path
func TestLoadEnglish(t *testing.T) {
	err := SetLanguage("en")

	assert.Nil(t, err)
	assert.Equal(t, "", NotError.Message, "Default valut must be: ")
	assert.Equal(t, "Error in the request parameters", ErrorMiddlewareQueryParams.Message, "Default valut must be: Error in the request parameters")

}

func TestLoadSpanish(t *testing.T) {
	err := SetLanguage("es")

	assert.Nil(t, err)
	assert.Equal(t, "", NotError.Message, "Default valut must be: ")
	assert.Equal(t, "Error en los parametros de la solicitud", ErrorMiddlewareQueryParams.Message, "Default valut must be: Error en los parametros de la solicitud")
}
