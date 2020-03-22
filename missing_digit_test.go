package missing_digit_exercise

import "testing"
import "github.com/stretchr/testify/assert"

func Test_SimpleAddition(t *testing.T) {
	res, _ := MissingDigit("2 + 4x = 44")
	assert.Equal(t, 2, res)
}

func Test_SecondAddition(t *testing.T) {
	res, _ := MissingDigit("x + 10 = 10")
	assert.Equal(t, 0, res)
}

func Test_ThirdAddition(t *testing.T) {
	res, _ := MissingDigit("2 + 12 = 1x")
	assert.Equal(t, 4, res)
}

func Test_FourthAddition(t *testing.T) {
	res, _ := MissingDigit("2x + 312 = 335")
	assert.Equal(t, 3, res)
}

func Test_FirstSubtraction(t *testing.T) {
	res, _ := MissingDigit("2x - 2 = 20")
	assert.Equal(t, 2, res)
}

func Test_SecondSubtraction(t *testing.T) {
	res, _ := MissingDigit("20 - x = 16")
	assert.Equal(t, 4, res)
}

func Test_ThirdSubtraction(t *testing.T) {
	res, _ := MissingDigit("20 - 3 = 1x")
	assert.Equal(t, 7, res)
}

func Test_FirstMultiplication(t *testing.T) {
	res, _ := MissingDigit("2x * 3 = 75")
	assert.Equal(t, 5, res)
}

func Test_SecondMultiplication(t *testing.T) {
	res, _ := MissingDigit("2 * x = 16")
	assert.Equal(t, 8, res)
}

func Test_ThirdMultiplication(t *testing.T) {
	res, _ := MissingDigit("3 * 3 = x")
	assert.Equal(t, 9, res)
}

func Test_FirstDivision(t *testing.T) {
	res, _ := MissingDigit("2x / 5 = 5")
	assert.Equal(t, 5, res)
}

func Test_SecondDivision(t *testing.T) {
	res, _ := MissingDigit("2 / x = 1")
	assert.Equal(t, 2, res)
}

func Test_ThirdDivision(t *testing.T) {
	res, _ := MissingDigit("9 / 3 = x")
	assert.Equal(t, 3, res)
}

func Test_WeirdOperator(t *testing.T) {
	_, err := MissingDigit("9 $ 3 = x")
	assert.Error(t, err)
}

func Test_WeirdFirstNumber(t *testing.T) {
	_, err := MissingDigit("a + 3 = x")
	assert.Error(t, err)
}

func Test_WeirdSecondNumber(t *testing.T) {
	_, err := MissingDigit("3 + a = x")
	assert.Error(t, err)
}

func Test_WeirdThirdNumber(t *testing.T) {
	_, err := MissingDigit("3 + 3 = ax")
	assert.Error(t, err)
}

func Test_MissingWhiteSpaces(t *testing.T) {
	_, err := MissingDigit("3+3=x")
	assert.Error(t, err)
}

func Test_UnsatisfiableEquation(t *testing.T) {
	_, err := MissingDigit("3 + 3 = 9x")
	assert.Error(t, err)
}
