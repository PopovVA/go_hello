package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {

	// Division by zero
	err := Validate(0, "/")
	assert.Error(t, err)

	// Unsupported operation
	err = Validate(0, "h")
	assert.Error(t, err)

	// Valid case
	err = Validate(10, "+")
	assert.Nil(t, err)

}
