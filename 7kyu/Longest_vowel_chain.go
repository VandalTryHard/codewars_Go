// 7 kyu
// Longest vowel chain
// The vowel substrings in the word codewarriors are o,e,a,io. The longest of these has a length of 2.
// Given a lowercase string that has alphabetic characters only (both vowels and consonants) and no spaces,
// return the length of the longest vowel substring. Vowels are any of aeiou.

// Good luck!

// If you like substring Katas, please try:

// Non-even substrings

// Vowel-consonant lexicon
// Fundamentals
// Strings

// package kata

// import "sort"

// func Solve(s string) int {
// 	check := "aeiou"
// 	vowel_string := ""
// 	vowel_list := []string{}

// 	for _, i := range s {
// 		if i == rune(check[0]) || i == rune(check[1]) || i == rune(check[2]) || i == rune(check[3]) || i == rune(check[4]) {
// 			vowel_string += string(i)
// 		} else {
// 			vowel_list = append(vowel_list, vowel_string)
// 			vowel_string = ""
// 		}
// 	}
// 	sort.Slice(vowel_list, func(i, j int) bool {
// 		return len(vowel_list[i]) > len(vowel_list[j])
// 	})
// 	return len(vowel_list[0])
// }

// How this works
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(Solve("chrononhotonthuooaos"))
}

func Solve(s string) int {
	check := "aeiou"
	vowel_string := ""
	vowel_list := []string{}

	for _, i := range s {
		if i == rune(check[0]) || i == rune(check[1]) || i == rune(check[2]) || i == rune(check[3]) || i == rune(check[4]) {
			vowel_string += string(i)
		} else {
			// fmt.Println(vowel_string)
			vowel_list = append(vowel_list, vowel_string)
			vowel_string = ""
		}
	}
	fmt.Println(vowel_list)
	sort.Slice(vowel_list, func(i, j int) bool {
		return len(vowel_list[i]) > len(vowel_list[j])
	})
	fmt.Println(vowel_list)
	return len(vowel_list[0])
}
