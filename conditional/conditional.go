package conditional

import "fmt"

// CheckNumber normal If
func CheckNumber(number float64) (result string) {
	if number > 0 {
		result = "It is a positive number"
	} else if number < 0 {
		result = "It is a negative number"
	} else {
		result = "It is zero"
	}
	return
}

// CheckLessOrGreater if with assignment
func CheckLessOrGreater(number float64) (result string) {
	if limit := 240.25; number > limit {
		result = "Greater"
	} else if number < limit {
		result = "Less"
	} else {
		result = "Equal"
	}

	return
}

// CheckSign idea way
// minimum branches and return early as possible
func CheckSign(number int) {
	if number % 2 == 0 {
		fmt.Println("Positive number")
		return
	}
	fmt.Println("Negative number")
}