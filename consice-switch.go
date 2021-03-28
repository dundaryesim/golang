package main

import "fmt"

func main() {
	fmt.Println("There is a cavern entrance here and path to east.")
	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("You head further up The mountain.")
	case "enter cave", "go inside":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that")

	}

}
