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

// Switch statement
func switch_example(day int) string {
	// No need for break statements; Go automatically breaks after each case
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

// Switch with multiple discrete conditions
func conditional_switch_example(score int) string {
	switch score {
	case 90, 100:
		return "A"
	case 80, 89:
		return "B"
	case 70, 79:
		return "C"
	default:
		return "F"
	}

}
