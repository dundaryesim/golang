package main

import "fmt"

func main() {
	var red uint8 = 255
	red++
	fmt.Println(red)
	//büyük unit tanımlamak başa dönmesini engeller.
	var number int16 = 127
	number++
	fmt.Println(number)
	// %08b sekiz bite tamamlar
	var green uint8 = 3
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)
}
