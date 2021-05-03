package main

import "fmt"

func main() {
	// Indices indicate number of bytes! not runes!!!
	spanishText := "¿Cómo estás?"

	ilkIkiHarf := spanishText[0:3] // ¿C, because ¿ holds 2 bytes

	fmt.Println(ilkIkiHarf)

	// Create a slice directly with a) var, b) type inference. There is an array behind!

	// array
	var x [3]string // ["" "" ""]

	// var olan bir elemani degistiriyorsun!
	// eleman eklemiyorsun!
	x[0] = "yesim"
	x[1] = "serhat"
	x[2] = "umut"

	y := [3]string{
		"yesim",
		"serhat",
		"umut",
	}

	fmt.Println(x, y)

	// slice
	var z []string // []

	// you can't do that!
	// index out of range [0] with length 0
	// because our slice doesn't have items to change!
	// z[0] = "yesim"
	// z[1] = "serhat"
	// z[2] = "umut"

	// instead, you should do this:
	z = append(z, "yesim")
	z = append(z, "serhat")
	z = append(z, "umut")

	// var olan bir elemani degistirmiyorsun! eleman ekliyorsun!

	t := []string{
		"yesim",
		"serhat",
		"umut",
	}

	fmt.Println(z, t)

}
