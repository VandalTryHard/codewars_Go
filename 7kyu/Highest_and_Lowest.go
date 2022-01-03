/*
7 kyu
Highest and Lowest

In this little assignment you are given a string of space separated numbers,
and have to return the highest and lowest number.
Examples

HighAndLow("1 2 3 4 5")  // return "5 1"
HighAndLow("1 2 -3 4 5") // return "5 -3"
HighAndLow("1 9 3 4 -5") // return "9 -5"

Notes

    All numbers are valid Int32, no need to validate them.
    There will always be at least one number in the input string.
    Output string must be two numbers separated by a single space, and highest number is first.

Fundamentals
Strings
*/

/*
package kata
import (
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	strs := strings.Split(in, " ")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}
	sort.Ints(ary)
	solution := strconv.Itoa(ary[len(ary)-1]) + " " + strconv.Itoa(ary[0])
	return solution
}
*/

// How this works
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
}

func HighAndLow(in string) string {
	strs := strings.Split(in, " ")
	ary := make([]int, len(strs))
	for i := range ary {
		ary[i], _ = strconv.Atoi(strs[i])
	}
	sort.Ints(ary)
	solution := strconv.Itoa(ary[len(ary)-1]) + " " + strconv.Itoa(ary[0])
	return solution
}
