package collections

// MapExample demonstrates the use of maps in Go
// Maps are collections of key-value pairs where keys are unique and you can retrieve values using their corresponding keys.
func MapExample() {
	// Creating a map with string keys and int values
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35

	// Accessing values in the map
	println("Alice's age is", ages["Alice"])
	println("Bob's age is", ages["Bob"])
	println("Charlie's age is", ages["Charlie"])
}

// Initalizing a map with a composite literal
func MapLiteralExample() {
	var personAge = map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}
	println("Alice's age is", personAge["Alice"])
	println("Bob's age is", personAge["Bob"])
	println("Charlie's age is", personAge["Charlie"])
}

// Checking if a key exists in a map
func MapKeyExistenceExample() {
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}
	age, exists := ages["Bob"] // Check if "Bob" exists in the map
	if exists {
		println("Bob's age is", age)
	} else {
		println("Bob's age not found")
	}
	// Note if you try to access a key that doesn't exist, it will return the default value of that data type.
	// In the case of an int, it will return zero value for the value type (0 for int, "" for string, etc.)
}

// Deleting a key-value pair from a map
func MapDeleteExample() {
	ages := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
	}
	println("Before deletion, Bob's age is", ages["Bob"])
	delete(ages, "Bob")                                  // Delete the key "Bob" from the map. First argument is the map, second argument is the key to delete.
	println("After deletion, Bob's age is", ages["Bob"]) // This will print 0 since Bob has been deleted
}
