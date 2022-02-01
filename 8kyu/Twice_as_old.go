/*
8 kyu
Twice as old

Your function takes two arguments:

    current father's age (years)
    current age of his son (years)

Сalculate how many years ago the father was twice as old as his son (or in how many years he will be twice as old).
Fundamentals
Mathematics
Algorithms
Numbers
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(TwiceAsOld(36, 7))
}

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	return int(math.Abs(float64(dadYearsOld - (sonYearsOld * 2))))
}
