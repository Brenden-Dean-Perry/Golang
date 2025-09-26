package main

import (
	"fmt"
	"golang/datatypes"
)

func main() {
	const pi = 3.14
	var x = 5

	fmt.Println("Hello, World!")
	datatypes.PrintTypes()
	fmt.Println()
	fmt.Println("2 + 3 =", Addit(2, 3))
	fmt.Println("5 * pi =", float64(x)*pi)
	fmt.Println()
}
