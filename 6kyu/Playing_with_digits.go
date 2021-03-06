/*
6 kyu
Playing with digits

Some numbers have funny properties. For example:

    89 --> 8¹ + 9² = 89 * 1

    695 --> 6² + 9³ + 5⁴= 1390 = 695 * 2

    46288 --> 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51

Given a positive integer n written as abcd... (a, b, c, d... being digits) and a positive integer p

    we want to find a positive integer k, if it exists, such as the sum of the digits of n taken to the successive powers of p is
	equal to k * n.

In other words:

    Is there an integer k such as : (a ^ p + b ^ (p+1) + c ^(p+2) + d ^ (p+3) + ...) = n * k

If it is the case we will return k, if not return -1.

Note: n and p will always be given as strictly positive integers.

digPow(89, 1) should return 1 since 8¹ + 9² = 89 = 89 * 1
digPow(92, 1) should return -1 since there is no k such as 9¹ + 2² equals 92 * k
digPow(695, 2) should return 2 since 6² + 9³ + 5⁴= 1390 = 695 * 2
digPow(46288, 3) should return 51 since 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51

Fundamentals
Mathematics
Algorithms
Numbers
*/

/*
package kata

import (
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	sum := 0.0
	for i, s := range strconv.Itoa(n) {
		sum += math.Pow(float64(s-'0'), float64(p+i))
	}
	intSum := int(sum)
	if intSum%n == 0 {
		return intSum / n
	}
	return -1
}
*/

// How this works
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(DigPow(89, 1))
	fmt.Println(DigPow(92, 1))
	fmt.Println(DigPow(695, 2))
	fmt.Println(DigPow(46288, 3))
	fmt.Println(DigPow(3263, 4))
}

func DigPow(n, p int) int {
	sum := 0.0
	for i, s := range strconv.Itoa(n) {
		sum += math.Pow(float64(s-'0'), float64(p+i))
	}
	intSum := int(sum)
	if intSum%n == 0 {
		return intSum / n
	}
	return -1
}
