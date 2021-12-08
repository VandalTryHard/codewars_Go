//8 kyu
// Abbreviate a Two Word Name
// Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.
// The output should be two capital letters with a dot separating them.
// It should look like this:
// Sam Harris => S.H
// patrick feeney => P.F

// package kata
// import (
// 	"strings"
// )
// func AbbrevName(name string) string{
//   split_name := strings.Split(name, " ")
// 	split_first := strings.Split(split_name[0], "")
// 	split_last := strings.Split(split_name[1], "")
//   return strings.Title(split_first[0]) + "." + strings.Title(split_last[0])
// }

// How this works
package main

import (
	"fmt"
	"strings"
)

func abbrevName(name string) {
	split_name := strings.Split(name, " ")
	split_first := strings.Split(split_name[0], "")
	split_last := strings.Split(split_name[1], "")

	fmt.Println(strings.Title(split_first[0]) + "." + strings.Title(split_last[0]))
}

func main() {
	abbrevName("Koko kola")
}
