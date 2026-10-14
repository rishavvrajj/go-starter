package mathutil

import (
	"fmt"
	"math"
)

// Pi is a constant exported from this package
const Pi = math.Pi

// package-level variable (exported)
var Version string = "1.0.0"

// unexported package-level variable
var internalCounter int

// init function runs automatically when package is imported
func init() {
	fmt.Println("mathutil package initialized")
	internalCounter = 0
}

// Add exports two numbers (exported function)
func Add(a, b float64) float64 {
	internalCounter++
	return a + b
}

// Subtract exports two numbers (exported function)
func Subtract(a, b float64) float64 {
	internalCounter++
	return a - b
}

// multiply is an unexported helper function (not accessible outside package)
func multiply(a, b float64) float64 {
	return a * b
}

// Multiply exports multiplication function
func Multiply(a, b float64) float64 {
	internalCounter++
	return multiply(a, b)
}

// Divide exports division function
func Divide(a, b float64) (float64, error) {
	internalCounter++
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// GetInternalCounter returns the internal counter (exported)
// This demonstrates accessing unexported variables through exported methods
func GetInternalCounter() int {
	return internalCounter
}

// ResetInternalCounter resets the internal counter (exported)
func ResetInternalCounter() {
	internalCounter = 0
}

// AreaOfCircle calculates area of a circle given radius
func AreaOfCircle(radius float64) float64 {
	return Pi * radius * radius
}

// Power calculates base raised to exponent
func Power(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}