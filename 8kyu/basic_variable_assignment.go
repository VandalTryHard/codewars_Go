/*8 kyu
Basic variable assignment
This code should store "codewa.rs" as a variable called name but it's not working. Can you figure out why?
Bugs
*/

// package kata
// func Namevar() string {
//   var a string = "code"
//   var b string = "wa.rs"
//   var name string = a + b
//   return name
// }

// How this works
package main

import "fmt"

func main() {
	var a string = "code"
	var b string = "wa.rs"
	var name string = a + b
	fmt.Println(name)
}
