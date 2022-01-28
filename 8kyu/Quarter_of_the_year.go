/*
8 kyu
Quarter of the year

Given a month as an integer from 1 to 12, return to which quarter of the year it belongs as an integer number.

For example: month 2 (February), is part of the first quarter; month 6 (June), is part of the second quarter; and month 11 (November),
 is part of the fourth quarter.
Fundamentals
*/

/*
package kata

func QuarterOf(month int) int {
	if month >= 1 && month <= 3 {
		return 1
	} else if month > 3 && month <= 6 {
		return 2
	} else if month > 6 && month <= 9 {
		return 3
	} else {
		return 4
	}
}
*/

// How this works
package main

import "fmt"

func main() {
	fmt.Println(QuarterOf(8))
}

func QuarterOf(month int) int {
	if month >= 1 && month <= 3 {
		return 1
	} else if month > 3 && month <= 6 {
		return 2
	} else if month > 6 && month <= 9 {
		return 3
	} else {
		return 4
	}
}
