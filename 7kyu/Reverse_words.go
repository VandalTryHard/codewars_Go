/*
7 kyu
Reverse words

Complete the function that accepts a string parameter, and reverses each word in the string.
All spaces in the string should be retained.
Examples

"This is an example!" ==> "sihT si na !elpmaxe"
"double  spaces"      ==> "elbuod  secaps"

Fundamentals
Strings
*/

/*
package kata

func ReverseWords(str string) string {
	var rev string
	var word string

	for _, i := range str {
		if i == ' ' {
			rev = rev + word + " " // Adds word and space to result
			word = ""              // Empties word variable
		} else {
			word = string(i) + word // Adds letter to temporary word variable
		}
	}

	return rev + word // reverse those words
}
*/

// How this works
package main

import (
	"fmt"
)

func main() {
	fmt.Println(ReverseWords("The quick brown fox jumps over the lazy dog."))
}

func ReverseWords(str string) string {
	var rev string
	var word string

	for _, i := range str {
		if i == ' ' {
			rev = rev + word + " " // Adds word and space to result
			word = ""              // Empties word variable
		} else {
			word = string(i) + word // Adds letter to temporary word variable
		}
	}

	return rev + word // reverse those words
}
