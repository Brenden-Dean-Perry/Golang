package main

var array = [3]string{} // 3 elements, all initialized to the zero value for strings ("")
// Arrays have a fixed size and type and stored in contiguous memory locations.
// Arrays don't need to know each memory address, just the starting address and the size of each element.

func index_indexing() {
	array[0] = "Learn Go"
	array[1] = "Build an app"
	array[2] = "Deploy to cloud"

	println("First task:", array[0])
	println("Second task:", array[1:3])
}

func immediately_initialized_array() {
	numbers := [5]int{1, 2, 3, 4, 5} // Array with 5 integers
	println("Numbers:", numbers)
}

func array_size_inference() {
	// The size of the array is inferred by the compiler from the number of initializers
	fruits := [...]string{"Apple", "Banana", "Cherry"} // Array with 3 strings.
	println("Fruits:", fruits)
}
