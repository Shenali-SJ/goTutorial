package conditional

import (
	"fmt"
	"math/rand"
)

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

func CheckDay() {
	day := 3
	fmt.Println("Day : ", day)

	//day2's scope is limited to switch block
	switch day2 := 34; day2 {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid day")
	}
}

func CheckVowels(letter string) {
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}
}

func number() int {
	//return 9 - 11
	return 9 + 11
}

func CheckValue() {
	num := number()
	switch {
	case num < 20:
		if num < 0 {
			fmt.Println("Less than zero")
			break
		}
		fmt.Println("Greater than 20")
		fallthrough
	case num < 30:
		fmt.Println("Greater than 30")
		fallthrough
	case num < 40:
		fmt.Println("Greater than 40")
	}
}

func CheckEven() {
	randLoop:
		for {
			switch i := rand.Intn(100); {
			case i % 2 == 0:
				fmt.Printf("Even number is generated %d", i)
				break randLoop
			}
		}
}