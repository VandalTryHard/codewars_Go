/*
7 kyu
Previous multiple of three

Given a positive integer n: 0 < n < 1e6, remove the last digit until you're left with a number
that is a multiple of three.

Return n if the input is already a multiple of three, and return -1 if no such number exists.
Examples

1      => -1
25     => -1
36     => 36
1244   => 12
952406 => 9

Fundamentals
*/

/*
package kata

func PrevMultOfThree(n int) interface{} {
	for i := n; i > 0; i /= 10 {
		if i%3 == 0 {
			return i
		}
	}
	return nil
}
*/

// How this works
package main

import "fmt"

func main() {
	fmt.Println(PrevMultOfThree(1244))
}

func PrevMultOfThree(n int) interface{} {

	for i := n; i > 0; i /= 10 {
		if i%3 == 0 {
			return i
		}
	}

	return nil
}
