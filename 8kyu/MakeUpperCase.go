/*
8 kyu
MakeUpperCase

Write a function which converts the input string to uppercase.
Fundamentals
*/

/*
package kata
import "strings"
func MakeUpperCase(str string) string {
	return strings.ToUpper(str)
}
*/

// How this works
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MakeUpperCase("hello world"))
}

func MakeUpperCase(str string) string {
	return strings.ToUpper(str)
}
