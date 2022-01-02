/*
6 kyu
Highest Rank Number in an Array

Complete the method which returns the number which is most frequent in the given input array.
If there is a tie for most frequent number, return the largest number among them.

Note: no empty arrays will be given.
Examples

[12, 10, 8, 12, 7, 6, 4, 10, 12]              -->  12
[12, 10, 8, 12, 7, 6, 4, 10, 12, 10]          -->  12
[12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10]  -->   3

Fundamentals
Arrays
Objects
*/

/*
package kata

func HighestRank(nums []int) int {
	mii, maxK, maxV := map[int]int{}, 0, 0
	for _, v := range nums {
		mii[v]++
		if mii[v] > maxV || (mii[v] == maxV && v > maxK) {
			maxK = v
			maxV = mii[v]
		}
	}
	return maxK
}
*/

// How this works
package main

import (
	"fmt"
)

func main() {
	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12}))
	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 10}))
}

func HighestRank(nums []int) int {
	mii, maxK, maxV := map[int]int{}, 0, 0
	for _, v := range nums {
		mii[v]++
		if mii[v] > maxV || (mii[v] == maxV && v > maxK) {
			maxK = v
			maxV = mii[v]
		}
	}
	return maxK
}
