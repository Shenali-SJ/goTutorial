package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Println("Hello World")

	fmt.Println()

	// declaring without initializing
	// default zero value will be assigned depending on the type
	var num1 int
	fmt.Println("Declaration only: ", num1)

	// declaring + initializing : with type
	var num2 int = 2345
	fmt.Println("Assignment: ", num2)

	// declaring + initializing : without type
	// type will be inferred depending on the value assigned
	var num3 = 454
	fmt.Println("Assignment without type: ", num3)

	// multiple variables (same type) - with type
	var height, weight int = 158, 43
	fmt.Println("Old height and weight: ", height, weight)

	weight = 55
	fmt.Println("New weight after dieting for 10 years: ", weight)

	// multiple variables (same type) - without type
	var id, index = 3244, 23
	fmt.Println("id: ", id, " / index: ", index)

	// multiple variables (different types)
	var (
		name = "Shenali"
		age = 20
		studentId int // if not assigning, should specify the type
	)
	fmt.Println("Name: ", name, ", age: ", age, ", Student ID: ", studentId)

	studentId = 20191152
	fmt.Println("Student ID: ", studentId)

	//shorthand
	number := 34
	fmt.Println(number)

	stName, stAge := "Shaini", 15
	fmt.Println("Student: ", stName, "(", stAge, ")")

	minNumber := math.Min(2323.2, 33545.0)
	fmt.Println("Minimum: ", minNumber)

	fmt.Println()

	// bool
	isCrazy, isStupid := true, false
	isComplicated := isCrazy && isStupid
	fmt.Println("Complicated? ", isComplicated)

	// int
	a := 34
	fmt.Printf("a : type is %T and size is %d \n", a, unsafe.Sizeof(a))
	//sizeOf returns the size in bytes
	// int - depending on the platform 32 or 64

	// float
	b := 34.2
	fmt.Printf("b : type is %T and size is %d \n", b, unsafe.Sizeof(b))

	// complex
	c1 := complex(3, 2)
	c2 := 5 + 2i

	cSum := c1 + c2
	fmt.Println(cSum)

	//type conversion - explicit
	var n1 int =  21
	var f1 float64 = float64(n1)
	fmt.Println(f1)

	fmt.Println()

	// constants
	// untyped
	const const1 = "Only constant in life is change"
	fmt.Println(const1)

	// typed
	const const2 string = "Shenali"
	fmt.Println(const2)

	// group of constants
	const (
		constNum = 232
		constName = "CHANGE"
	)

	fmt.Println(constName)

	// custom types and constants - string
	const name1 = "Rachel"
	type customString string
	var name2 customString = "Monica"

	fmt.Println(name2)
	name2 = customString(name1)
	fmt.Println(name2)

	// custom types and constants - bool
	const bool1 = false
	type customBool bool
	var bool2 customBool = true

	fmt.Println(bool1, bool2)
	//bool1 = bool2 // type casting not allowed

	// custom types and constants - numeric
	const mysteryNumber = 12
	var mysteryNumberInt32 int32 = mysteryNumber
	var mysteryNumberInt64 int64 = mysteryNumber
	var mysteryNumberFloat64 float64 = mysteryNumber
	var mysteryNumberComplex64 complex64 = mysteryNumber

	fmt.Printf("mysteryNumber - %T, mysteryNumberInt32 %T, mysteryNumberInt64 %T, mysteryNumberFloat64 %T, mysteryNumberComplex64 %T", mysteryNumber, mysteryNumberInt32, mysteryNumberInt64, mysteryNumberFloat64, mysteryNumberComplex64)

}
