package main

// and_or_not_example demonstrates logical AND, OR, and NOT operations
func and_or_not_example(a bool, b bool) (bool, bool, bool) {
	// Logical AND
	andResult := a && b
	// Logical OR
	orResult := a || b
	// Logical NOT
	notA := !a
	return andResult, orResult, notA
}
