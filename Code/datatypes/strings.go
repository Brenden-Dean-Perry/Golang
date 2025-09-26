package datatypes

import (
	"fmt"
	"unicode/utf8"
)

// It is important to understand that strings in Go are a sequence of bytes, not characters.
// This distinction is crucial when dealing with multi-byte characters, such as those found in UTF-8 encoded text.
// For example, the string "résumé" contains characters that are represented by more than one byte each.
//  When you access a string by index, you are accessing the individual bytes, not the characters.
// This can lead to unexpected results if you are not careful.

// Additionally, it is important to understand utf8. A rune is a Unicode code point, which can be one or more bytes.
// UTF-8 is a variable-width encoding that uses one to four bytes to represent a rune.

func Indexing_a_string() {
	var myString string = "résumé"
	var indexed byte = myString[0] // Accessing the first byte
	fmt.Println("First byte (in bytes):", indexed)
}

// string_iteration demonstrates iterating over a string and printing byte indices and values
// Note: That iterating over a string with range iterates over runes, not bytes.
// So for multi-byte characters, the index may jump according to the number of bytes in the rune.
// However, the value will always be the rune itself (not the byte value at that index).
// In summary, Go has made it easy to work with strings and Unicode.
// But it's important to understand the underlying concepts of bytes, runes, and UTF-8 encoding to avoid pitfalls when manipulating strings.
func String_iteration() {
	var myString string = "résumé"
	fmt.Println("Iterating over bytes:")
	for i, v := range myString {
		fmt.Printf("Byte index: %d, Byte value: %v\n", i, v) // v is the byte value
	}
}

// string_iteration_runes demonstrates iterating over a string converted to a slice of runes and printing rune indices and values.
// This is useful when you want to work with individual characters (runes) in a string, especially when dealing with multi-byte characters.
// The length of the slice will be equal to the number of runes in the string, not the number of bytes.
// Thus, the number of runes will be equal to the length of the slice of the array as you would expect in other programming languages.
func String_iteration_runes() {
	var myString = []rune("résumé")
	fmt.Println("Iterating over runes:")
	for i, v := range myString {
		fmt.Printf("Rune index: %d, Rune value: %v\n", i, v) // v is the rune value
	}
}

// get_string_length_in_bytes returns the length of the string in bytes
func Get_string_length_in_bytes(s string) int {
	return len(s) // returns the length in bytes
}

// get_string_rune_count returns the number of runes (Unicode code points) in the string
func Get_string_rune_count(s string) int {
	return utf8.RuneCountInString(s) // returns the number of runes (Unicode code points)
}
