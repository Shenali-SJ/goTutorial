package simpleInterest

// Calculate exported function, first letter should be capital
func Calculate(p, r , t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}