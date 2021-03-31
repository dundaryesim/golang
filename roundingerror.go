package main

import "fmt"

func main() {
	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, "o F\n")
	fmt.Print((9.0/5.0*celsius)+32, "o F\n")
	// Doğrusu aşağıda
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Print(fahrenheit, "o F")
}
