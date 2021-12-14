// 8 kyu
// If you can't sleep, just count sheep!!
// If you can't sleep, just count sheep!!
// Task:

// Given a non-negative integer, 3 for example, return a string with a murmur: "1 sheep...2 sheep...3 sheep...". Input will always be valid, i.e. no negative integers.
// Fundamentals
// Loops
// Control Flow
// Basic Language Features

// package kata

// import (
// 	"fmt"
// 	"strings"
// )

// func countSheep(num int) string {
// 	str_ := []string{}
// 	for i := 1; i <= num; i++ {
// 		str_1 := fmt.Sprint(i) + " sheep..."
// 		str_ = append(str_, str_1)
// 	}
// 	return strings.Join(str_, "")
// }

// How this works
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countSheep(5))
}

func countSheep(num int) string {
	str_ := []string{}
	for i := 1; i <= num; i++ {
		str_1 := fmt.Sprint(i) + " sheep..."
		str_ = append(str_, str_1)
	}
	return strings.Join(str_, "")
}
