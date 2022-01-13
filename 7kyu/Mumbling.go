/*
7 kyu
Mumbling

This time no story, no theory. The examples below show you how to write function accum:
Examples:

accum("abcd") -> "A-Bb-Ccc-Dddd"
accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
accum("cwAt") -> "C-Ww-Aaa-Tttt"

The parameter of accum is a string which includes only letters from a..z and A..Z.
Fundamentals
*/

/*
package kata

import (
	"strings"
	"unicode/utf8"
)

func Accum(s string) string {
	acc := make([]string, utf8.RuneCount([]byte(s)))
	for i, j := range s {
		n := strings.ToUpper(string(j))
		if i > 0 {
			n += strings.Repeat(strings.ToLower(string(j)), i)
		}
		acc[i] = n
	}
	return strings.Join(acc, "-")
}
*/

// How this works
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(Accum("ZpglnRxqenU"))
}

func Accum(s string) string {
	acc := make([]string, utf8.RuneCount([]byte(s)))
	for i, j := range s {
		n := strings.ToUpper(string(j))
		if i > 0 {
			n += strings.Repeat(strings.ToLower(string(j)), i)
		}
		acc[i] = n
	}
	return strings.Join(acc, "-")
}
