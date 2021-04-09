package main

import "fmt"

func main() {
	str1 := "hello!"
	str1 = "world"

	// 119. Prints the code point value
	fmt.Println(str1[0])

	// "w". Prints the character
	fmt.Printf("%c\n", str1[0])

	// foo := str1[0:4] // 0'dan 4'e kadar al. veya [:4]
	// bar := str1[3:]  // 3'ten sona kadar
	sonKarakter := str1[4]
	fmt.Printf("%c\n", sonKarakter)

	// strings can't be altered
	// strings are immutable!
	// sonKarakter = "a"
}
