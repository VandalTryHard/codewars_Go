// 7 kyu
// Will you survive the zombie onslaught?
// I'm afraid you're in a rather unfortunate situation. You've injured your leg, and are unable
// to walk, and a number of zombies are shuffling towards you, intent on eating your brains.
// Luckily, you're a crack shot, and have your trusty rifle to hand.

// The zombies start at range metres, and move at 0.5 metres per second. Each second, you first shoot
// one zombie, and then the remaining zombies shamble forwards another 0.5 metres.

// If any zombies manage to get to 0 metres, you get eaten. If you run out of ammo before shooting all
//  the zombies, you'll also get eaten. To keep things simple, we can ignore any time spent reloading.

// Write a function that accepts the total number of zombies, a range in metres, and the number of bullets you have.

// If you shoot all the zombies, return "You shot all X zombies." If you get eaten before killing all the
// zombies, and before running out of ammo, return "You shot X zombies before being eaten: overwhelmed."
// If you run out of ammo before shooting all the zombies, return "You shot X zombies before being eaten: ran out of ammo."

// (If you run out of ammo at the same time as the remaining zombies reach you, return
// 	"You shot X zombies before being eaten: overwhelmed.".)

// Good luck! (I think you're going to need it.)
// Fundamentals
// Games
package main

import "fmt"

func Zombie_shootout(zombies, initial_range, ammo int) string {
	if zombies == 0 {
		return fmt.Sprintf("You shot all 0 zombies.")
	}

	var c int
	j := float64(initial_range)
	for i := j; i > .0; i = i - 0.5 {
		diff := ammo - 1
		if diff >= 0 {
			ammo = ammo - 1
			zombies = zombies - 1
			c++
		} else {
			break
		}
		if zombies == 0 {
			return fmt.Sprintf("You shot all %d zombies.", c)
		}
		j = j - 0.5
	}
	var t string
	if ammo == 0 && zombies >= 1 && j == 0.0 {
		t = "overwhelmed."
	} else if ammo == 0 {
		t = "ran out of ammo."
	} else {
		t = "overwhelmed."
	}
	return fmt.Sprintf("You shot %d zombies before being eaten: %s", c, t)
}

func main() {
	fmt.Println(Zombie_shootout(3, 10, 10))
	fmt.Println(Zombie_shootout(100, 8, 200))
	fmt.Println(Zombie_shootout(50, 10, 8))
}
