// 8 kyu
// Grasshopper - Messi goals function
// Messi goals function
// Messi is a soccer player with goals in three leagues:
//     LaLiga
//     Copa del Rey
//     Champions
// Complete the function to return his total number of goals in all three leagues.
// Note: the input will always be valid.
// For example:
// 5, 10, 2  -->  17
// Fundamentals

// package kata

// func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
//   sumGoals := laLigaGoals + copaDelReyGoals + championsLeagueGoals
// 	return sumGoals
// }

// How this works
package main

import "fmt"

func Goals(laLigaGoals, copaDelReyGoals, championsLeagueGoals int) int {
	sumGoals := laLigaGoals + copaDelReyGoals + championsLeagueGoals
	return sumGoals
}

func main() {
	fmt.Println(Goals(1, 2, 3))
}
