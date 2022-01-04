/*
6 kyu
Create Phone Number

Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of
those numbers in the form of a phone number.
Example

CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})  // returns "(123) 456-7890"

The returned format must be correct in order to complete this challenge.
Don't forget the space after the closing parentheses!
Algorithms
Arrays
Strings
Loops
Control Flow
Basic Language Features
Fundamentals
Formatting
Regular Expressions
Declarative Programming
Advanced Language Features
*/

/*
package kata

import "strconv"

func CreatePhoneNumber(numbers [10]uint) string {
	uintToString := ""
	for _, num := range numbers {
		uintToString += strconv.FormatUint(uint64(num), 10)
	}
	numberString := "(" + uintToString[0:3] + ") " + uintToString[3:6] + "-" + uintToString[6:]
	return numberString
}
*/

// How this works
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func CreatePhoneNumber(numbers [10]uint) string {
	uintToString := ""
	for _, num := range numbers {
		uintToString += strconv.FormatUint(uint64(num), 10)
	}
	numberString := "(" + uintToString[0:3] + ") " + uintToString[3:6] + "-" + uintToString[6:]
	return numberString
}
