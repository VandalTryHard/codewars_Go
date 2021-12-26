// 7 kyu
// Is this a triangle?
// Implement a function that accepts 3 integer values a, b, c. The function should return true if a
// triangle can be built with the sides of given length and false in any other case.

// (In this case, all triangles must have surface greater than 0 to be accepted).
// Fundamentals
// Mathematics
// Algorithms
// Numbers
// Utilities

// package kata

// func IsTriangle(a, b, c int) bool {
// 	var result bool
// 	if a+b > c && b+c > a && a+c > b {
// 		result = true
// 	} else {
// 		result = false
// 	}
// 	return result
// }

// How this works
package main

import "fmt"

func main() {
	fmt.Println(IsTriangle(5, 1, 2))
	fmt.Println(IsTriangle(4, 2, 3))
	fmt.Println(IsTriangle(5, 1, 5))
	fmt.Println(IsTriangle(-1, 2, 3))
}

func IsTriangle(a, b, c int) bool {
	var result bool
	if a+b > c && b+c > a && a+c > b {
		result = true
	} else {
		result = false
	}
	return result
}
