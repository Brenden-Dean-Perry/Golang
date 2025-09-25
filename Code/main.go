package main

import (
	"fmt"
)

func main() {
	const pi = 3.14
	var x = 5

	fmt.Println("Hello, World!")
	print_types()
	fmt.Println()
	fmt.Println("2 + 3 =", addit(2, 3))
	fmt.Println("5 * pi =", float64(x)*pi)
	fmt.Println()
}
