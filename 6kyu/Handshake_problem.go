/*
6 kyu
Handshake problem
Johnny is a farmer and he annually holds a beet farmers convention "Drop the beet".

Every year he takes photos of farmers handshaking. Johnny knows that no two farmers handshake more than once.
He also knows that some of the possible handshake combinations may not happen.

However, Johnny would like to know the minimal amount of people that participated this year just by counting
all the handshakes.

Help Johnny by writing a function, that takes the amount of handshakes and returns the minimal amount of people
needed to perform these handshakes (a pair of farmers handshake only once).
Algorithms
*/

/*
package kata

func GetParticipants(handshakes int) int {
	var people int = 1
	for i := 0; i <= handshakes; i++ {
		if handshakes > (people*(people-1))/2 {
			people += 1
		}
	}
	return people
}
*/

// How this works and look this https://nrich.maths.org/6708/solution
package main

import "fmt"

func main() {
	fmt.Println(GetParticipants(0))
	fmt.Println(GetParticipants(1))
	fmt.Println(GetParticipants(3))
	fmt.Println(GetParticipants(6))
	fmt.Println(GetParticipants(7))

}

func GetParticipants(handshakes int) int {
	var people int = 1
	for i := 0; i <= handshakes; i++ {
		if handshakes > (people*(people-1))/2 {
			people += 1
		}
	}
	return people
}
