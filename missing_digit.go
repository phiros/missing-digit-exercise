package missing_digit_exercise

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type equationChecker func(num1, num2, num3 int) bool

// MissingDigit takes as input an equation as a string which
// contains three numbers in total and one operand.
// An operand may be +, -, * or /.
// One number in the equation must contain an "x" (a missing digit).
// The function returns this missing digit as a int.
// Example:
// MissingDigit("2x + 4 = 28") will return 4
//
// Note that all tokens in the equation must be separated by a
// whitespace.
func MissingDigit(s string) (int, error) {
	tokens := strings.Split(s, " ")
	if len(tokens) != 5 {
		return 0, fmt.Errorf("wrong format for equation: \"%s\" (missing whitespace?)", s)
	}
	op := tokens[1]

	switch op {
	case "+":
		i, err := findDigit(tokens, func(num1, num2, num3 int) bool {
			return num1+num2 == num3
		})
		return i, err
	case "-":
		i, err := findDigit(tokens, func(num1, num2, num3 int) bool {
			return num1-num2 == num3
		})
		return i, err
	case "*":
		i, err := findDigit(tokens, func(num1, num2, num3 int) bool {
			return num1*num2 == num3
		})
		return i, err
	case "/":
		i, err := findDigit(tokens, func(num1, num2, num3 int) bool {
			if num2 == 0 {
				return false
			}
			return num1/num2 == num3
		})
		return i, err
	}

	return 0, fmt.Errorf("cannot handle: \"%s\" please check equation format", s)
}

func findDigit(tokens []string, isEquation equationChecker) (int, error) {
	for i := 0; i < 10; i++ {
		firstNum, secondNum, thirdNum, err := numbersWithFilledDigit(tokens, i)
		if err != nil {
			return 0, err
		}
		if isEquation(firstNum, secondNum, thirdNum) {
			return i, nil
		}
	}
	return 0, errors.New("cannot determine missing digit")
}

func numbersWithFilledDigit(tokens []string, missingDigit int) (int, int, int, error) {
	firstNumAsString := tokens[0]
	secondNumAsString := tokens[2]
	thirdNumAsString := tokens[4]
	firstNumAsString = replaceXWithDigit(firstNumAsString, missingDigit)
	secondNumAsString = replaceXWithDigit(secondNumAsString, missingDigit)
	thirdNumAsString = replaceXWithDigit(thirdNumAsString, missingDigit)
	firstNum, err := strconv.Atoi(firstNumAsString)
	if err != nil {
		return 0, 0, 0, err
	}
	secondNum, err := strconv.Atoi(secondNumAsString)
	if err != nil {
		return 0, 0, 0, err
	}
	thirdNum, err := strconv.Atoi(thirdNumAsString)
	if err != nil {
		return 0, 0, 0, err
	}
	return firstNum, secondNum, thirdNum, nil
}

func replaceXWithDigit(numberAsString string, missingDigit int) string {
	if strings.Contains(numberAsString, "x") {
		numberAsString = strings.Replace(numberAsString, "x", strconv.Itoa(missingDigit), 1)
	}
	return numberAsString
}
