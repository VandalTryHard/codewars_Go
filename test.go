package main

import (
	"fmt"
	"strings"
)

func main() {
	// PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz")
	fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
}

func PrinterError(s string) string {
	s := strings.Split(s, "")
	// check := strings.Split("abcdefghijklm", "")
	// col := 0
	// for i, s := range s {
	// 	if s == check {
	// 		&col = *col + 1
	// 	}
	// }
	// result = "%v/%v", col, len(s)
	// result := strings.Join(s, ",")
	return s
}
