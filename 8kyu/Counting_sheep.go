// 8 kyu
// Counting sheep...
// Consider an array/list of sheep where some sheep may be missing from their place.
//We need a function that counts the number of sheep present in the array (true means present).
// For example,
// [true,  true,  true,  false,
//   true,  true,  true,  true ,
//   true,  false, true,  false,
//   true,  false, false, true ,
//   true,  true,  true,  true ,
//   false, false, true,  true]
// The correct answer would be 17.
// Hint: Don't forget to check for bad values like null/undefined
// Fundamentals
// Arrays

// package kata

// func CountSheeps(numbers []bool) int {
//   result := 0
//   for i := 0; i<len(numbers); i++{
//     if numbers[0+i] == true{
//       result = result +1
//     }
//   }
//   return result
// }

// How this works
package main

import "fmt"

func CountSheeps(arr1 []bool) int {
	result := 0
	for i := 0; i < len(arr1); i++ {
		if arr1[0+i] == true {
			result = result + 1
		}
	}
	return result
}

func main() {
	arr1 := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}
	fmt.Println(CountSheeps(arr1))
}
