package main

import (
	"fmt"
)

func main() {
	check := "ATCG"
	fmt.Println(rune(check[0]))
	fmt.Println(DNAStrand("ATCG"))
	fmt.Println(DNAStrand("ATTGC"))
	fmt.Println(DNAStrand("GTAT"))
}

func DNAStrand(dna string) string {
	chek := "ATCG"
	new_ := ""
	for _, s := range dna {
		switch s {
		case rune(chek[0]):
			new_ += "T"
		case rune(chek[1]):
			new_ += "A"
		case rune(chek[2]):
			new_ += "G"
		case rune(chek[3]):
			new_ += "C"
		}
	}
	return new_
}
