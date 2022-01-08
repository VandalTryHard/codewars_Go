/*
7 kyu
Simple remove duplicates

Remove the duplicates from a list of integers, keeping the last ( rightmost ) occurrence of each element.
Example:

For input: [3, 4, 4, 3, 6, 3]

    remove the 3 at index 0
    remove the 4 at index 1
    remove the 3 at index 3

Expected output: [4, 6, 3]

More examples can be found in the test cases.

Good luck!
Fundamentals
*/

/*
package kata

func Solve(arr []int) (res []int) {
	visited := map[int]bool{}
	for i := len(arr) - 1; i >= 0; i-- {
		n := arr[i]
		if visited[n] {
			continue
		}
		visited[n] = true
		res = append([]int{n}, res...)
	}
	return
}
*/

// How this works
package main

import "fmt"

func main() {
	fmt.Println(Solve([]int{3, 4, 4, 3, 6, 3}))
	fmt.Println(Solve([]int{1, 2, 1, 2, 1, 2, 3}))
}

func Solve(arr []int) (res []int) {
	visited := map[int]bool{}
	for i := len(arr) - 1; i >= 0; i-- {
		n := arr[i]
		if visited[n] {
			continue
		}
		visited[n] = true
		res = append([]int{n}, res...)
	}
	return
}
