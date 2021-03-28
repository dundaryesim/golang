//Adventure games are divided up into rooms.
//Write a program that uses ifand else if to display a description for each of three rooms: cave, entrance, and mountain.
//When writing your program, ensure the curly braces {} are placed according to the one true brace style as shown in listing 3.3.

package main

import "fmt"

func main() {
	var room = "cave"
	if room == "cave" {
		fmt.Println("You will be in the cave")
	} else if room == "entrance" {
		fmt.Println("You will be enterance of the cave")
	} else if room == "mountain" {
		fmt.Println("You must find the cave")
	} else {
		fmt.Println("You are nowhere")
	}
}
