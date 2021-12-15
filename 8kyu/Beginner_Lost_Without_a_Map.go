// 8 kyu
// Beginner - Lost Without a Map

// Given an array of integers, return a new array with each value doubled.

// For example:

// [1, 2, 3] --> [2, 4, 6]
// Fundamentals
// Arrays
// Numbers

// package kata
// func Maps(x []int) []int {
// 	n := []int{}
// 	for _, value := range x {
// 		n = append(n, value*2)
// 	}
// 	return n
// }

// How this works
package main

import "fmt"

func main() {
	x := []int{4, 1, 1, 1, 4}
	fmt.Println(Maps(x))
}

func Maps(x []int) []int {
	n := []int{}
	for _, value := range x {
		n = append(n, value*2)
	}
	return n
}
