package main

import "fmt"

func main() {
	// code smell! duplicate-repeating code!
	// var serhat struct {
	// 	isim string
	// 	yas  int
	// }

	// var yesim struct {
	// 	isim string
	// 	yas  int
	// }

	// var umut struct {
	// 	isim string
	// 	yas  int
	// }

	type person struct {
		isim string
		yas  int
	}

	var serhat person
	serhat.isim = "Serhat"
	serhat.yas = 30

	// composite literal
	umut := person{isim: "Umut", yas: 3}

	fmt.Println(serhat.isim, umut.isim)
}
