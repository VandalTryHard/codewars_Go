/*
6 kyu
Highest Scoring Word

Given a string of words, you need to find the highest scoring word.

Each letter of a word scores points according to its position in the alphabet: a = 1, b = 2, c = 3 etc.

You need to return the highest scoring word as a string.

If two words score the same, return the word that appears earliest in the original string.

All letters will be lowercase and all inputs will be valid.
Fundamentals
Strings
Arrays
Numbers
*/

/*
package kata

import "strings"

func wordScore(s string) (scor byte) {
  for i := range s {
    scor += s[i] - 96
  }
  return
}

func High(s string) string {
  var scor, scorNew byte
  var word string
  for _, wd := range strings.Split(s, " ") {
    scorNew = wordScore(wd)
    if scorNew > scor {
      scor = scorNew
      word = wd
    }
  }
  return word
}
*/

// How this works
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(High("man i need a taxi up to ubud"))
}

func wordScore(s string) (scor byte) {
	for i := range s {
		scor += s[i] - 96
	}
	return
}

func High(s string) string {
	var scor, scorNew byte
	var word string
	for _, wd := range strings.Split(s, " ") {
		scorNew = wordScore(wd)
		if scorNew > scor {
			scor = scorNew
			word = wd
		}
	}
	return word
}
