package main

import "fmt"

func main() {
	// anonim fonksiyonu tanimla ve degiskene ata.
	// ama henuz cagirma!
	square := func(x int) int {
		return x * x
	}

	// anonim fonksiyonu cagir ve ekrana bastir
	fmt.Println(square(10))

	// anonim fonksiyonu tanimla, hemen cagir ve dondugu degeri degiskene ata
	x := func(x int) int {
		return x * x
	}(10)

	// degiskene atanmis degeri ekrana bas
	fmt.Println(x)
}
