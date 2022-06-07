package calculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleCalculate() {
	res, _ := Calculate(10, 10, "+")
	fmt.Println(res)
	// Output: 20
}

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
