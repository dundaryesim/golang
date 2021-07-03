package main

import "fmt"

func main() {
	// Çoklu yazım
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(len(planets)) // Len built-in function
	println(planets[7])       // You access the first element of the planets array

}
