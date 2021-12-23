package main

import (
	"fmt"
	"math"
)

func main() {
	// kek := math.Ceil(52 / 100)
	// kek := 52.0 / 100.0
	// fmt.Println(math.Ceil(kek))
	var l, w, h float64 = 6.3, 4.5, 3.29
	roll := (52.0 / 100.0) * 10.0
	fmt.Println(roll)
	room := math.Ceil(((l + w) * h) * 2)
	fmt.Println(room)
	resultRoll := (room / roll) * 1.15
	fmt.Println(resultRoll)
	rolls := math.Ceil(resultRoll)
	fmt.Println(rolls)
	// result := numbers[rolls]
	// fmt.Println(result)
}
