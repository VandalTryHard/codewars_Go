/*
8 kyu
Remove First and Last Character

It's pretty straightforward. Your goal is to create a function that removes the first and last
characters of a string. You're given one parameter, the original string. You don't have to worry
with strings with less than two characters.
Fundamentals
Basic Language Features
Strings
*/

/*
package kata

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
*/

// How this works
package main

import "fmt"

func main() {
	fmt.Println(RemoveChar("country"))
}

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
