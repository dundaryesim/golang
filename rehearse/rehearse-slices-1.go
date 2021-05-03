package main

import "fmt"

func main() {
	// modifying the original array will also modify slices created from it
	// modifying slices taken from an array, will also modify the original array
	myArr := [5]string{
		"yesim",
		"serhat",
		"umut",
		"cakil",
		"yoldas",
	}

	fmt.Printf("Dizi => %v, turu => %[1]T\n", myArr)

	// half open range
	insanlar := myArr[:2]
	hayvanlar := myArr[2:]

	fmt.Printf("Insanlar => %v, turu => %[1]T\n", insanlar)
	fmt.Printf("Hayvanlar => %v, turu => %[1]T\n", hayvanlar)

	kopekler := hayvanlar[1:]
	kuslar := hayvanlar[:1]

	fmt.Printf("Kopekler => %v, turu => %[1]T\n", kopekler)
	fmt.Printf("Kuslar => %v, turu => %[1]T\n", kuslar)

	myArr[0] = "ozgur"

	fmt.Println("Ilk myArr elemani Ozgur ile degistirildi!!!")
	fmt.Println("")

	fmt.Printf("Dizi => %v, turu => %[1]T\n", myArr)
	fmt.Printf("Insanlar => %v, turu => %[1]T\n", insanlar)
	fmt.Printf("Hayvanlar => %v, turu => %[1]T\n", hayvanlar)
	fmt.Printf("Kopekler => %v, turu => %[1]T\n", kopekler)
	fmt.Printf("Kuslar => %v, turu => %[1]T\n", kuslar)

	kopekler[1] = "cavcav"

	fmt.Println("Ilk kopekler elemani cavcav ile degistirildi!!!")
	fmt.Println("")

	fmt.Printf("Dizi => %v, turu => %[1]T\n", myArr)
	fmt.Printf("Insanlar => %v, turu => %[1]T\n", insanlar)
	fmt.Printf("Hayvanlar => %v, turu => %[1]T\n", hayvanlar)
	fmt.Printf("Kopekler => %v, turu => %[1]T\n", kopekler)
	fmt.Printf("Kuslar => %v, turu => %[1]T\n", kuslar)

	// Slice an entire array with full open range
	mySlice := myArr[:]

	fmt.Printf("myArr'nin turu => %T, mySlice'in turu => %T\n", myArr, mySlice)
}
