/*
8 kyu
Reversed sequence
Build a function that returns an array of integers from n to 1 where n>0.

Example : n=5 --> [5,4,3,2,1]
Fundamentals
*/

/*
package kata

func ReverseSeq(n int) []int {
	var x []int
	for i := n; i > 0; i-- {
		x = append(x, i)
	}
	return x
}
*/

// How this works
package main

import "fmt"

func main() {
	fmt.Println(ReverseSeq(5))
}

func ReverseSeq(n int) []int {
	var x []int
	for i := n; i > 0; i-- {
		x = append(x, i)
	}
	return x
}
