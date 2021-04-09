package main

import (
	"fmt"
)

func main() {
	// len is a built-in function

	isim1 := "yesim"
	isim2 := "yeşim"

	// fmt.Println(len(isim1), len(isim2)) // 5, 6: len counts bytes!

	// for i := 0; i < len(isim1); i++ {
	// 	fmt.Printf("%c => %[1]v\n", isim1[i]+1)
	// }

	for i := 0; i < len(isim1); i++ {
		fmt.Printf("%c => %[1]v\n", isim1[i])
	}

	fmt.Println("===========")

	// string[index] works with ASCII (uint8-byte)
	// range works with utf8 (int32-rune)
	for index, harf := range isim2 {
		c := harf + 1
		fmt.Printf("Index numarasi %v ve harf degeri %c\n", index, c)
	}
}
