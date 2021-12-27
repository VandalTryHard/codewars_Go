// 7 kyu
// String ends with?
// Complete the solution so that it returns true if the first argument(string)
// passed in ends with the 2nd argument (also a string).

// Examples:

// solution("abc", "bc") // returns true
// solution("abc", "d") // returns false

// Fundamentals
// Strings

/*
package kata

import "strings"

func solution(str, ending string) bool {
	var result bool
	if strings.Contains(str, ending) {
		result = true
	} else {
		result = false
	}
	return result
}
*/

// How this works
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solution("abc", "c"))
	fmt.Println(solution("$a = $b + 1", "+1"))
	fmt.Println(solution("abc", "c"))
	fmt.Println(solution("samurai", "ra"))
}

func solution(str, ending string) bool {
	var result bool
	if strings.Contains(str, ending) {
		result = true
	} else {
		result = false
	}
	return result
}
