package main

// If-else statement
func if_else_example(x int) string {
	if x%2 == 0 {
		return "even"
	} else if x < 0 {
		return "negative"
	} else {
		return "odd"
	}
}
