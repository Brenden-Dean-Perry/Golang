package datatypes

import (
	"fmt"
)

// PrintTypes prints various data types and their values
func PrintTypes() {
	var a int = 42
	var b int8 = 8
	var c int16 = 16
	var d int32 = 32
	var e int64 = 64
	var f uint = 42
	var g uint8 = 8
	var h uint16 = 16
	var i uint32 = 32
	var j uint64 = 64
	var k float64 = 3.14
	var l bool = true
	var m string = "Hello, Go!"
	var n rune = 'G'
	var o byte = 255
	var p complex128 = complex(1, 2)

	fmt.Printf("a: %v, type: %T\n", a, a)
	fmt.Printf("b: %v, type: %T\n", b, b)
	fmt.Printf("c: %v, type: %T\n", c, c)
	fmt.Printf("d: %v, type: %T\n", d, d)
	fmt.Printf("e: %v, type: %T\n", e, e)
	fmt.Printf("f: %v, type: %T\n", f, f)
	fmt.Printf("g: %v, type: %T\n", g, g)
	fmt.Printf("h: %v, type: %T\n", h, h)
	fmt.Printf("i: %v, type: %T\n", i, i)
	fmt.Printf("j: %v, type: %T\n", j, j)
	fmt.Printf("k: %v, type: %T\n", k, k)
	fmt.Printf("l: %v, type: %T\n", l, l)
	fmt.Printf("m: %v, type: %T\n", m, m)
	fmt.Printf("n: %v, type: %T\n", n, n)
	fmt.Printf("o: %v, type: %T\n", o, o)
	fmt.Printf("p: %v, type: %T\n", p, p)
}
