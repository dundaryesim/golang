package main

import "fmt"

func main() {
	// arrays can store only a single type!

	// declare an array with var
	var dizi [5]int

	// prints zero values!
	// [0 0 0 0 0]
	fmt.Println(dizi)

	// assign some values to array
	dizi[0] = 30
	dizi[1] = 12

	// [30 12 0 0 0]
	fmt.Println(dizi)

	// 5
	fmt.Println(len(dizi))

	// 30
	fmt.Println(dizi[0])

	// invalid array index 6 (out of bounds for 5-element array)
	// fmt.Println(dizi[6])

	// composite literal. hem olustur hem degerleri ata!
	anotherDizi := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(anotherDizi)

	// more readable way
	someDizi := [5]int{
		1,
		9,
		5,
		4,
		7,
	}
	fmt.Println(someDizi)

	// let go count items for you
	smartDizi := [...]string{
		"serhat",
		"ferhat",
		"merhat",
		"celat",
		"kerhat",
		"verhat",
	}

	fmt.Println(smartDizi)

	for i := 0; i < len(smartDizi); i++ {
		fmt.Println(smartDizi[i])
	}

	for sira, item := range smartDizi {
		fmt.Println(sira, "====>", item)
	}

	ilkAile := [3]string{"serhat", "yesim", "umut"}

	// diziyi yeni bir degiskene atamak kopyasini olusturur!
	yeniAile := ilkAile
	yeniAile[2] = "stevie!"

	// [serhat yesim umut]
	fmt.Println(ilkAile)

	// [serhat yesim stevie!]
	fmt.Println(yeniAile)
}
