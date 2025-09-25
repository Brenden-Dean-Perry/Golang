package main

import "unicode/utf8"

// get_string_length_in_bytes returns the length of the string in bytes
func get_string_length_in_bytes(s string) int {
	return len(s) // returns the length in bytes
}

// get_string_rune_count returns the number of runes (Unicode code points) in the string
func get_string_rune_count(s string) int {
	return utf8.RuneCountInString(s) // returns the number of runes (Unicode code points)
}
