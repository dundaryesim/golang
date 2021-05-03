package main

import "fmt"

func main() {
	// strings are immutable!
	isim := "Yesim"

	ilkHarf := isim[0:1]

	ilkHarf = "S"

	fmt.Println(ilkHarf, isim)
}
