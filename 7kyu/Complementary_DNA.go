// 7 kyu
// Complementary DNA
// Deoxyribonucleic acid (DNA) is a chemical found in the nucleus of cells and carries the
// "instructions" for the development and functioning of living organisms.

// If you want to know more: http://en.wikipedia.org/wiki/DNA

// In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G".
// You have function with one side of the DNA (string, except for Haskell); you need to get
// the other complementary side. DNA strand is never empty or there is no DNA at all
// (again, except for Haskell).

// More similar exercise are found here: http://rosalind.info/problems/list-view/ (source)

// Example: (input --> output)

// "ATTGC" --> "TAACG"
// "GTAT" --> "CATA"

// dnaStrand []        `shouldBe` []
// dnaStrand [A,T,G,C] `shouldBe` [T,A,C,G]
// dnaStrand [G,T,A,T] `shouldBe` [C,A,T,A]
// dnaStrand [A,A,A,A] `shouldBe` [T,T,T,T]

// Fundamentals
// Strings

// package kata

// func DNAStrand(dna string) string {
// 	chek := "ATCG"
// 	new_ := ""
// 	for _, s := range dna {
// 		switch s {
// 		case rune(chek[0]):
// 			new_ += "T"
// 		case rune(chek[1]):
// 			new_ += "A"
// 		case rune(chek[2]):
// 			new_ += "G"
// 		case rune(chek[3]):
// 			new_ += "C"
// 		}
// 	}
// 	return new_
// }

// How this works
package main

import (
	"fmt"
)

func main() {
	fmt.Println(DNAStrand("ATCG"))
	fmt.Println(DNAStrand("ATTGC"))
	fmt.Println(DNAStrand("GTAT"))
}

func DNAStrand(dna string) string {
	chek := "ATCG"
	new_ := ""
	for i := 0; i < len(dna); i++ {
		if dna[i] == chek[0] {
			new_ += "T"
		} else if dna[i] == chek[1] {
			new_ += "A"
		} else if dna[i] == chek[2] {
			new_ += "G"
		} else if dna[i] == chek[3] {
			new_ += "C"
		}
	}
	return new_
}

// or

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(DNAStrand("ATCG"))
// 	fmt.Println(DNAStrand("ATTGC"))
// 	fmt.Println(DNAStrand("GTAT"))
// }

// func DNAStrand(dna string) string {
// 	chek := "ATCG"
// 	new_ := ""
// 	for _, s := range dna {
// 		switch s {
// 		case rune(chek[0]):
// 			new_ += "T"
// 		case rune(chek[1]):
// 			new_ += "A"
// 		case rune(chek[2]):
// 			new_ += "G"
// 		case rune(chek[3]):
// 			new_ += "C"
// 		}
// 	}
// 	return new_
// }
