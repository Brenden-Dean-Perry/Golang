package main

import "fmt"

// deferExample demonstrates the use of defer to delay execution of a function until the surrounding function returns
func deferExample() {
	defer fmt.Println("world") // This will be executed last
	fmt.Println("hello")       // This will be executed first
}
