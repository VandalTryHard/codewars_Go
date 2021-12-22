// 7 kyu
// Printer Errors
// In a factory a printer prints labels for boxes. For one kind of boxes the printer has to use colors which, for the sake
// of simplicity, are named with letters from a to m.

// The colors used by the printer are recorded in a control string. For example a "good" control string would be
// aaabbbbhaijjjm meaning that the printer used three times color a, four times color b, one time color h then one time
// color a...

// Sometimes there are problems: lack of colors, technical malfunction and a "bad" control string is produced e.g.
// aaaxbbbbyyhwawiwjjjwwm with letters not from a to m.

// You have to write a function printer_error which given a string will return the error rate of the printer as a string
// representing a rational whose numerator is the number of errors and the denominator the length of the control string.
// Don't reduce this fraction to a simpler expression.

// The string has a length greater or equal to one and contains only letters from ato z.
// Examples:

// s="aaabbbbhaijjjm"
// printer_error(s) => "0/14"

// s="aaaxbbbbyyhwawiwjjjwwm"
// printer_error(s) => "8/22"

// Fundamentals

// package kata

// import (
// 	"strconv"
// 	"strings"
// )

// func PrinterError(s string) string {
// 	mas_s := strings.Split(s, "")
// 	col := 0
// 	for _, val := range mas_s {
// 		if val != "a" && val != "b" && val != "c" && val != "d" && val != "e" && val != "f" && val != "g" && val != "h" && val != "i" && val != "j" && val != "k" && val != "l" && val != "m" {
// 			col = col + 1
// 		}
// 	}
// 	result := strconv.Itoa(col) + "/" + strconv.Itoa(len(mas_s))
// 	return result
// }

// How thiw works
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
}

func PrinterError(s string) string {
	mas_s := strings.Split(s, "")
	fmt.Println(mas_s)
	check := strings.Split("abcdefghijklm", "")
	fmt.Println(check)
	col := 0
	fmt.Println(col)
	for _, val := range mas_s {
		if val != "a" && val != "b" && val != "c" && val != "d" && val != "e" && val != "f" && val != "g" && val != "h" && val != "i" && val != "j" && val != "k" && val != "l" && val != "m" {
			col = col + 1
		}
	}
	result := strconv.Itoa(col) + "/" + strconv.Itoa(len(mas_s))
	return result
}
