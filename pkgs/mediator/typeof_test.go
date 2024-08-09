package mediator_test

import (
	"api.ddd/pkgs/mediator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeOf(t *testing.T) {

	// Arrange
	expectedTypeOf := "TestStruct"

	// Act
	typeOf := mediator.TypeOf(&TestStruct{})

	// Assert
	assert.NotNil(t, typeOf)
	assert.Equal(t, expectedTypeOf, typeOf)
}

func TestTypeOfNilShouldReturnEmptyString(t *testing.T) {

	// Act
	typeOf := mediator.TypeOf(nil)

	// Assert
	assert.NotNil(t, typeOf)
	assert.Empty(t, typeOf)
}
