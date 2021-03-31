package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nickel := 0.05
	dime := 0.10
	quarter := 0.25

	moneys := []float64{nickel, dime, quarter}

	for balance := 0.0; balance < 20.0; {
		balance += moneys[rand.Intn(3)]
		fmt.Printf("Kasadaki para: $%5.2f\n", balance)
	}
}
