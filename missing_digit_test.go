package missing_digit_exercise

import "testing"
import "github.com/stretchr/testify/assert"

func Test_SimpleAddition(t *testing.T) {
	res := MissingDigit("2 + 4x = 44")
	assert.Equal(t, 2, res)
}

func Test_SecondAddition(t *testing.T) {
	res := MissingDigit("x + 10 = 10")
	assert.Equal(t, 0, res)
}
