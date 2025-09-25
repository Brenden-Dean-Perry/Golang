package main

import (
	"errors"
)

// Simple addition function
func addit(a int, b int) int {
	return a + b
}

// Returns multiple values
func return_multiple_values(a int, b int) (int, int) {
	return a + b, a * b
}

// Division with error handling
func division(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
