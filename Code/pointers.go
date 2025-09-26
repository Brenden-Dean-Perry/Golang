package main

import "fmt"

func Declare_pointer_variable() {
	var p *int32 = new(int32) // pointer to an int32 (aka a memory address that holds an int32 value).
	// Note we need to use the asterisk to denote a pointer type.
	// Also, we use the new function to allocate memory for the int32 value and return a pointer to it.
	// If we just declared var p *int32, it would be a nil pointer and would cause a runtime error if we tried to dereference it (aka access the value at that memory address).

	fmt.Println("Pointer address:", p) // Print the memory address stored in the pointer
	fmt.Println("Pointer value:", *p)  // Dereference the pointer to get the value at that address (should be 0, the zero value for int32)
}

// set_value_at_pointer sets the value at the memory address pointed to by the pointer
func Set_value_at_pointer() {
	var p *int32 = new(int32) // pointer to an int32
	*p = 42                   // Set the value at the memory address to 42
	// Note we use the asterisk to dereference the pointer and set the value at that address.

	fmt.Println("Pointer address:", p)
	fmt.Println("Pointer value:", *p)
}

// nil_pointer_example demonstrates a nil pointer (which does not point to any memory address)
func Nil_pointer_example() {
	var p *int32 // nil pointer (does not point to any memory address)
	// Note that if we try to dereference a nil pointer, it will cause a runtime error (panic).
	if p == nil {
		fmt.Println("Pointer is nil")
	} else {
		fmt.Println("Pointer address:", p)
		fmt.Println("Pointer value:", *p) // This would cause a runtime error if p is nil
	}
}

// address_of_variable demonstrates getting the address of a variable using the address-of operator (&)
func Address_of_variable() {
	var p *int32 = new(int32) // pointer to an int32
	var x int32
	p = &x  // Set the pointer to the address of the variable x using the address-of operator (&)
	*p = 42 // Set the value at the memory address to 42

	fmt.Println("Variable address:", p)
	fmt.Println("Variable value:", *p)
	fmt.Println("x value:", x) // x should now be 42 since both p and x point to the same memory address and we set the value at that address to 42
}

// passing_by_reference demonstrates passing an array by reference using a pointer.
// This allows the function to modify the original array rather than a copy of the array.
func Passing_by_reference(data *[5]int32) {
	// Modify the array elements
	for i := 0; i < len(data); i++ {
		data[i] *= 2
	}
}
