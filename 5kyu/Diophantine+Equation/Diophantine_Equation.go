/*
5 kyu
Diophantine Equation

In mathematics, a Diophantine equation is a polynomial equation, usually with two or more unknowns, such that only the integer
solutions are sought or studied.

In this kata we want to find all integers x, y (x >= 0, y >= 0) solutions of a diophantine equation of the form:
x2 - 4 * y2 = n

(where the unknowns are x and y, and n is a given positive number) in decreasing order of the positive xi.

If there is no solution return [] or "[]" or "". (See "RUN SAMPLE TESTS" for examples of returns).
Examples:

solEquaStr(90005) --> "[[45003, 22501], [9003, 4499], [981, 467], [309, 37]]"
solEquaStr(90002) --> "[]"

Hint:

x2 - 4 * y2 = (x - 2*y) * (x + 2*y)
Fundamentals
Mathematics
Algorithms
Numbers
*/

/*
package kata

import (
  "math"
)

func Solequa(n int) [][]int {
  r := make([][]int, 0)
  for i:= 1; i < int(math.Sqrt(float64(n))) + 1; i++ {
      if n % i == 0 {
          d := n/i;
          if (d - i) % 4 == 0 {
              y := (d - i) / 4;
              x := i + 2*y;
              r = append(r, []int{x, y})
          }
      }
  }
  return r;
}
*/

//How this works
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Solequa(9005))
}

func Solequa(n int) [][]int {
	// your code
	r := make([][]int, 0)
	for i := 1; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			d := n / i
			if (d-i)%4 == 0 {
				y := (d - i) / 4
				x := i + 2*y
				r = append(r, []int{x, y})
			}
		}
	}
	return r
}
