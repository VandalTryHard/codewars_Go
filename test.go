package main

import "fmt"

func main() {
	fmt.Println(inviteMoreWomen([]int{1, -1, 1}))
}

func inviteMoreWomen(L []int) bool {
	result := 0
	for _, num := range L {
		result += num
	}
	return result > 0
}
