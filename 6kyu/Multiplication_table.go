/*
6 kyu
Multiplication table

Your task, is to create NxN multiplication table, of size provided in parameter.

for example, when given size is 3:

1 2 3
2 4 6
3 6 9

for given example, the return value should be: [[1,2,3],[2,4,6],[3,6,9]]
Fundamentals
Arithmetic
Mathematics
Algorithms
Numbers
Arrays
*/

/*
package multiplicationtable

func MultiplicationTable(size int) [][]int {
	columns := [][]int{}
	for i := 1; i <= size; i++ {
		rows := []int{}
		for j := 1; j <= size; j++ {
			rows = append(rows, i*j)
		}
		columns = append(columns, rows)
	}
	return columns
}
*/

// How this works
package main

import "fmt"

func main() {
	fmt.Println(MultiplicationTable(3))
}

func MultiplicationTable(size int) [][]int {
	columns := [][]int{}
	for i := 1; i <= size; i++ {
		rows := []int{}
		for j := 1; j <= size; j++ {
			rows = append(rows, i*j)
		}
		columns = append(columns, rows)
	}
	return columns
}
