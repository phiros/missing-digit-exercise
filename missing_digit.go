package missing_digit_exercise

import (
	"strconv"
	"strings"
)

func MissingDigit(s string) int {
	tokens := strings.Split(s, " ")

	for i := 0; i < 10; i++ {
		firstNumAsString := tokens[0]
		secondNumAsString := tokens[2]
		thirdNumAsString := tokens[4]
		if strings.Contains(firstNumAsString, "x") {
			firstNumAsString = strings.Replace(firstNumAsString, "x", strconv.Itoa(i), 1)
		}
		if strings.Contains(secondNumAsString, "x") {
			secondNumAsString = strings.Replace(secondNumAsString, "x", strconv.Itoa(i), 1)
		}
		if strings.Contains(thirdNumAsString, "x") {
			thirdNumAsString = strings.Replace(thirdNumAsString, "x", strconv.Itoa(i), 1)
		}

		firstNum, _ := strconv.Atoi(firstNumAsString)
		secondNum, _ := strconv.Atoi(secondNumAsString)
		thirdNum, _ := strconv.Atoi(thirdNumAsString)

		if firstNum+secondNum == thirdNum {
			return i
		}
	}

	return 0
}
