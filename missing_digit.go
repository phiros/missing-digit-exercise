package missing_digit_exercise

import (
	"strconv"
	"strings"
)

type equationChecker func(num1, num2, num3 int) bool

func MissingDigit(s string) int {
	tokens := strings.Split(s, " ")
	op := tokens[1]

	switch op {
	case "+":
		i := findDigit(tokens, func(num1, num2, num3 int) bool {
			return num1+num2 == num3
		})
		return i
	case "-":
		i := findDigit(tokens, func(num1, num2, num3 int) bool {
			return num1-num2 == num3
		})
		return i
	case "*":
		i := findDigit(tokens, func(num1, num2, num3 int) bool {
			return num1*num2 == num3
		})
		return i
	case "/":
		i := findDigit(tokens, func(num1, num2, num3 int) bool {
			if num2 == 0 {
				return false
			}
			return num1/num2 == num3
		})
		return i
	}

	return 0
}

func findDigit(tokens []string, isEquation equationChecker) int {
	for i := 0; i < 10; i++ {
		firstNum, secondNum, thirdNum := NumbersWithFilledDigit(tokens, i)
		if isEquation(firstNum, secondNum, thirdNum) {
			return i
		}
	}
	return 0
}

func NumbersWithFilledDigit(tokens []string, missingDigit int) (int, int, int) {
	firstNumAsString := tokens[0]
	secondNumAsString := tokens[2]
	thirdNumAsString := tokens[4]
	firstNumAsString = replaceXWithDigit(firstNumAsString, missingDigit)
	secondNumAsString = replaceXWithDigit(secondNumAsString, missingDigit)
	thirdNumAsString = replaceXWithDigit(thirdNumAsString, missingDigit)
	firstNum, _ := strconv.Atoi(firstNumAsString)
	secondNum, _ := strconv.Atoi(secondNumAsString)
	thirdNum, _ := strconv.Atoi(thirdNumAsString)
	return firstNum, secondNum, thirdNum
}

func replaceXWithDigit(numberAsString string, missingDigit int) string {
	if strings.Contains(numberAsString, "x") {
		numberAsString = strings.Replace(numberAsString, "x", strconv.Itoa(missingDigit), 1)
	}
	return numberAsString
}
