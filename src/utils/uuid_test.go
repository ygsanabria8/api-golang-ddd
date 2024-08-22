package utils_test

import (
	"api.ddd/src/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhenCallGetUUIDShouldReturnUUIDString(t *testing.T) {
	// Act
	uuid := utils.GetUUID()

	// Assert
	assert.NotNil(t, uuid)
	assert.NotEmpty(t, uuid)
}

func TestGivenValidUUIDWhenCAllIsValidUUIDShouldReturnUUIDString(t *testing.T) {
	// Arrange
	uuid := utils.GetUUID()

	// Act
	uuidObject, err := utils.IsValidUUID(uuid)

	// Assert
	assert.NotNil(t, uuidObject)
	assert.NotEmpty(t, uuidObject)
	assert.Nil(t, err)
}

func TestGivenInvalidUUIDWhenCAllIsValidUUIDShouldReturnError(t *testing.T) {
	// Arrange
	uuid := "adalsdjasd"

	// Act
	_, err := utils.IsValidUUID(uuid)

	// Assert
	assert.NotNil(t, err)
}
