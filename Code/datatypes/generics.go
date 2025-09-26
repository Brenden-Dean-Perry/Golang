package datatypes

// demo of generics in Go
// Generic functions allow you to write functions that can operate on different types without sacrificing type safety.
// This is achieved using type parameters, which are specified in square brackets [] after the function name.
// sumSlice takes a slice of numbers (int, float32, or float64) and returns their sum
func SumSlice[T int | float32 | float64](numbers []T) T {
	var sum T
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Any allows any type, similar to interface{} in earlier Go versions
// isEmpty checks if a slice of any type is empty
func IsEmpty[T any](numbers []T) bool {
	return len(numbers) == 0
}

// Generics can also be used with structs
// Pair is a generic struct that holds two values of the same type
type Pair[T any] struct {
	First  T
	Second T
}

type gasEngine struct {
	gallons uint
	mpg     uint
}

type electricEngine struct {
	batteryCapacity uint // in kWh
	rangeMiles      uint // range in miles
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T // can be gasEngine or electricEngine
}
