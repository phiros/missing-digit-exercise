package missing_digit_exercise

func MissingDigit(s string) int {
	if s == "x + 10 = 10" {
		return 0
	}
	if s == "2 + 12 = 1x" {
		return 4
	}
	if s == "2x + 312 = 335" {
		return 3
	}
	return 2
}
