package main

func Constants_and_variables() {
	const pi = 3.14 // constant declaration
	var x = 5       // variable declaration with type inference
	println("5 * pi =", float64(x)*pi)
}

func Delay_variable_declaration() {
	var a int // declaration without initialization, defaults to zero value
	a = 10
	println("a:", a)
}

func Short_variable_declaration() {
	b := 20 // shorthand for var b int = 20. Type is inferred.
	println("b:", b)
}

func Multiple_variable_declaration() {
	var x, y, z int = 1, 2, 3
	var m, n, o = 4, "hello", 5.6 // type inference
	p, q, r := 7, "world", 8.9    // shorthand declaration with type inference

	println("x:", x, "y:", y, "z:", z)
	println("m:", m, "n:", n, "o:", o)
	println("p:", p, "q:", q, "r:", r)
}
