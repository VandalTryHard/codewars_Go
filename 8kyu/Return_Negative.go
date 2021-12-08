// 8 kyu
// Return Negative
// In this simple assignment you are given a number and have to make it negative. But maybe the number is already negative?
// Examples
// MakeNegative(1)    # return -1
// MakeNegative(-5)   # return -5
// MakeNegative(0)    # return 0
// Notes
//     The number can be negative already, in which case no change is required.
//     Zero (0) is not checked for any specific sign. Negative zeros make no mathematical sense.
// Fundamentals
// Numbers

// package kata

// func MakeNegative(x int) int {
//   if x < 0 {
// 		return x
// 	}
// 	return -1 * x
// }

// How this works
package main

import "fmt"

func MakeNegative(x int) int {
	if x < 0 {
		return x
	}
	return -1 * x
}

func main() {
	fmt.Println(MakeNegative(-1), MakeNegative(1), MakeNegative(0))
}
