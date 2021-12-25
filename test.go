package main

import (
	"fmt"
)

func main() {
	fmt.Println(Potatoes(99, 100, 98))
}

func Potatoes(p0, w0, p1 int) int {
	w1 := w0 * (100 - p0) / (100 - p1)
	return w1
}
