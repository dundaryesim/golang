package main

import "fmt"

type fiyat int

func (f fiyat) vergi() int {
	return int((f * 18 / 100) + f)
}

func main() {
	var laptopFiyat fiyat = 250

	fmt.Println(laptopFiyat.vergi())
}
