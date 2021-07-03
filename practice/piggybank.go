package main

import (
	"fmt"
	"math/rand"
)

func main() {

	nickels := 5
	dimes := 10
	quarters := 25

	for total := 0; total < 2000; {
		switch rand.Intn(3) {
		case 0:
			total += nickels
		case 1:
			total += dimes
		case 2:
			total += quarters
		}
		dollars := total / 100
		cents := total % 100
		fmt.Printf("$%d.%02d\n", dollars, cents)

	}
}
