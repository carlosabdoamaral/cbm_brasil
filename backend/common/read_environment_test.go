package common

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetEnvVariablesSuccess(t *testing.T) {
	errEnv := godotenv.Load("../test.env")
	assert.Nil(t, errEnv, nil)

	testVar := os.Getenv("VAR_FILL")
	assert.NotEmpty(t, testVar)
	assert.Equal(t, "test", testVar)
}

func TestGetEnvVariablesEmptyError(t *testing.T) {
	errEnv := godotenv.Load("../test.env")
	assert.Nil(t, errEnv, nil)

	testVar := os.Getenv("VAR_EMPTY")
	assert.Empty(t, testVar)
}

func TestGetEnvVariablesNoFileError(t *testing.T) {
	errEnv := godotenv.Load("../not_found.env")
	assert.NotNil(t, errEnv, nil)
}
