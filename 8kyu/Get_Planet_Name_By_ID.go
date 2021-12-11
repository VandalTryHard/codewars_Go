// 8 kyu
// Get Planet Name By ID
// The function is not returning the correct values. Can you figure out why?

// Example (Input --> Output ):

// 3 --> "Earth"

// Bugs
// Control Flow
// Basic Language Features
// Fundamentals

// package kata

// func GetPlanetName(ID int) string {
//   	switch ID {
//     case 1:
//       return "Mercury"
//     case 2:
//       return "Venus"
//     case 3:
//       return "Earth"
//     case 4:
//       return "Mars"
//     case 5:
//       return "Jupiter"
//     case 6:
//       return "Saturn"
//     case 7:
//       return "Uranus"
//     case 8:
//     	return "Neptune"
//     }
//   return "Pluto"
// }

// How this works
package main

import "fmt"

func main() {
	fmt.Println(GetPlanet(1))
}

func GetPlanet(ID int) string {
	switch ID {
	case 1:
		return "Mercury"
	case 2:
		return "Venus"
	}
	return "Yep"
}
