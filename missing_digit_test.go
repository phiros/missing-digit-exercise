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

func Test_ThirdAddition(t *testing.T) {
	res := MissingDigit("2 + 12 = 1x")
	assert.Equal(t, 4, res)
}

func Test_FourthAddition(t *testing.T) {
	res := MissingDigit("2x + 312 = 335")
	assert.Equal(t, 3, res)
}

func Test_FirstSubtraction(t *testing.T) {
	res := MissingDigit("2x - 2 = 20")
	assert.Equal(t, 2, res)
}

func Test_SecondSubtraction(t *testing.T) {
	res := MissingDigit("20 - x = 16")
	assert.Equal(t, 4, res)
}

func Test_ThirdSubtraction(t *testing.T) {
	res := MissingDigit("20 - 3 = 1x")
	assert.Equal(t, 7, res)
}
