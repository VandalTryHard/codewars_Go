/*
7 kyu
Sort Numbers
Finish the solution so that it sorts the passed in array of numbers. If the function passes
in an empty array or null/nil value then it should return an empty array.

For example:

solution(c(1, 2, 3, 10, 5)) # should return c(1, 2, 3, 5, 10)
solution(NULL)              # should return NULL

Fundamentals
*/

/*
package kata

import "sort"

func SortNumbers(numbers []int) []int {
  sort.Ints(numbers)
  return numbers
}
*/

// How this works
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortNumbers([]int{1, 2, 10, 50, 5}))
	// SortNumbers([]int{1, 2, 10, 50, 5})
}

func SortNumbers(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}
