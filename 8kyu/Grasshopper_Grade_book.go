/*
8 kyu
Grasshopper - Grade book

Grade book

Complete the function so that it finds the average of the three scores passed to it and returns the letter value
associated with that grade.
Numerical Score 	Letter Grade
90 <= score <= 100 	'A'
80 <= score < 90 	'B'
70 <= score < 80 	'C'
60 <= score < 70 	'D'
0 <= score < 60 	'F'

Tested values are all between 0 and 100. Theres is no need to check for negative values or values greater than 100.
Fundamentals
*/

/*
package kata

func GetGrade(a, b, c int) rune {
	switch result := (a + b + c) / 3; {
	case result < 60:
		return 'F'
	case result < 70:
		return 'D'
	case result < 80:
		return 'C'
	case result < 90:
		return 'B'
	default:
		return 'A'
	}
}
*/

// How this works
package main

import "fmt"

func main() {
	fmt.Println(GetGrade(70, 70, 100))
}

func GetGrade(a, b, c int) rune {
	switch result := (a + b + c) / 3; {
	case result < 60:
		return 'F'
	case result < 70:
		return 'D'
	case result < 80:
		return 'C'
	case result < 90:
		return 'B'
	default:
		return 'A'
	}
}
