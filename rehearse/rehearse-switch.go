package main

import "fmt"

func main() {
	isim := "yesim"

	switch isim {
	case "serhat":
		fmt.Println("Not yesim!")
	case "kuran":
		fmt.Println("Not yesim again!")
	case "dundar":
		fmt.Println("yes!")
	default:
		fmt.Println("none of them!")
	}
}
