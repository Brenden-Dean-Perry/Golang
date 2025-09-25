package main

func for_loop_example() {
	// for loop with initialization, condition, and post statement
	for i := 0; i < 5; i++ {
		println("Iteration:", i)
	}
}

func while_loop_example() {
	// while loop using for with only a condition
	j := 0
	for j < 5 {
		println("While Loop Iteration:", j)
		j++
	}
}

func while_loop_with_break() {
	// while loop with break statement
	k := 0
	for {
		if k >= 5 {
			break
		}
		println("While Loop with Break Iteration:", k)
		k++
	}
}

func loop_with_continue() {
	// for loop with continue statement
	for m := 0; m < 10; m++ {
		if m%2 == 0 {
			continue
		}
		println("Loop with Continue Iteration:", m)
	}
}

func loop_with_range() {
	// for loop using range to iterate over a slice
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		println("Index:", index, "Value:", value)
	}
}

// Note: The order of iteration over maps is not specified and can vary from one iteration to the next.
func loop_over_map() {
	// for loop using range to iterate over a map
	person := map[string]string{"name": "Alice", "city": "Wonderland"}
	for key, value := range person {
		println("Key:", key, "Value:", value)
	}
}
