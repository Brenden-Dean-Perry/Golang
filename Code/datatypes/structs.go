package datatypes

import "fmt"

// Define a struct
type Person struct {
	Name    string
	Age     int
	Address Address
}

// Define a struct
type Address struct {
	Street string
	City   string
	Zip    string
}

// Adding methods to structs
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

// Function to create and return a struct instance
func build_structs() {
	// Create an instance of the struct
	p := Person{Name: "Alice", Age: 30, Address: Address{"123 Main St", "Anytown", "12345"}} // Using composite literal
	p2 := Person{}
	p2.Name = "Bob"
	p2.Age = 25
	p2.Address = Address{"456 Elm St", "Shelbyville", "67890"}

	p3 := Person{"Charlie", 35, Address{"789 Oak St", "Springfield", "12345"}} // Positional initialization

	// Call the methods on the struct instances
	p.Greet()
	p2.Greet()
	p3.Greet()
}
