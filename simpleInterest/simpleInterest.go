package simpleInterest

import "fmt"

func init () {
	fmt.Println("### Simple Interest package: initialized ###")
}

// Calculate calculates simple interest based on principle, rate and time
// exported function, first letter should be capital
func Calculate(p, r , t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}