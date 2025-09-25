package main

// slices are just a wrapper around arrays. They are more flexible and powerful than arrays.
// They have additional functionality and can be resized dynamically.

// To define a slice, we declare an array without specifying its size.
var slice = []string{} // An empty slice of strings

// adding elements to a slice
func add_elements_to_slice() {
	slice = append(slice, "Learn Go") // append adds elements to the end of the slice
	slice = append(slice, "Build an app")
	slice = append(slice, "Deploy to cloud")
}

// You computer will automatically resize the slice as needed. However there are performance implications to consider.
func slice_capacity() {
	var intSlice []int = []int{1, 2, 3} // An empty slice of integers with 3 initial elements
	intSlice = append(intSlice, 4)      // Append an element to the slice
	println("Length:", len(intSlice))   // should equal 4
	println("Capacity:", cap(intSlice)) // Capacity will be 6.
	// The growth strategy for the new capacity is dynamic and depends on the current capacity:
	// For smaller slices (capacity < 1024 elements): The capacity is typically doubled.
	// For larger slices (capacity >= 1024 elements): The capacity increases by a factor of approximately 1.25 (or 25%).
}
