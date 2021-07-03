package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		var number = 11
		var n = rand.Intn(100) + 1
		if n < number {
			fmt.Printf("%v is too small \n", n)
		} else if n > number {
			fmt.Printf("%v is too big \n", n)
		} else {
			fmt.Println("You got it")
		}
		break

	}
}
