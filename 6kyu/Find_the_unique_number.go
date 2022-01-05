/*
6 kyu
Find the unique number

There is an array with some numbers. All numbers are equal except for one. Try to find it!

findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55

It’s guaranteed that array contains at least 3 numbers.

The tests contain some very huge arrays, so think about performance.

This is the first kata in series:

    Find the unique number (this kata)
    Find the unique string
    Find The Unique

Fundamentals
Algorithms
Numbers
Arrays
*/

/*
package kata

func FindUniq(arr []float32) float32 {
	counter := map[float32]float32{}
	for _, item := range arr {
		counter[item]++
	}

	resp := make([]float32, 0, len(counter))
	for key, value := range counter {
		if value == 1 {
			resp = append(resp, key)
		}
	}
	return resp[0]
}
*/

// Hiw this works
package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}))
	fmt.Println(FindUniq([]float32{0, 0, 0.55, 0, 0}))
}

func FindUniq(arr []float32) float32 {
	counter := map[float32]float32{}
	for _, item := range arr {
		counter[item]++
	}

	resp := make([]float32, 0, len(counter))
	for key, value := range counter {
		if value == 1 {
			resp = append(resp, key)
		}
	}
	return resp[0]
}
