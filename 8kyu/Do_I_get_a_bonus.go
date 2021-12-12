// 8 kyu
// Do I get a bonus?
// It's bonus time in the big city! The fatcats are rubbing their paws in anticipation...
// but who is going to make the most money?

// Build a function that takes in two arguments (salary, bonus). Salary will be an integer, and bonus a boolean.

// If bonus is true, the salary should be multiplied by 10. If bonus is false, the fatcat
// did not make enough money and must receive only his stated salary.

// Return the total figure the individual will receive as a string prefixed with "£"
// (= "\u00A3", JS, Go, Java and Julia), "$" (C#, C++, Ruby, Clojure, Elixir, PHP, Python, Haskell and Lua)
// or "¥" (Rust).
// Algorithms
// Booleans
// Strings
// Numbers

// package kata

// import (
// 	"strconv"
// )

// func BonusTime(salary int, bonus bool) string {
// 	if bonus {
// 		total := salary * 10
// 		return "\u00A3"+strconv.Itoa(total)
// 	} else {
// 		total := salary
// 		return "\u00A3"+strconv.Itoa(total)
// 	}
// }

// How this works
package main

import (
	"fmt"
	"strconv"
)

func BonusTime(salary int, bonus bool) string {
	if bonus {
		total := salary * 10
		return "\u00A3" + strconv.Itoa(total)
	} else {
		total := salary
		return "\u00A3" + strconv.Itoa(total)
	}
}

func main() {
	fmt.Println(BonusTime(100, false))
}
