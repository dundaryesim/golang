package main

import "fmt"

func main() {
	// name1 := "yesim"
	// var name2 = "yesim"
	// var name3 string = "yesim"

	// Code point of space is 32
	fmt.Printf("%v\n", ' ')

	// raw string!
	fmt.Println(`
		this is a
		raw string
		and it goes to
		multiple lines`)

	// code point of smiley face
	fmt.Printf("%c\n", 65)
	fmt.Printf("%c\n", 128515)

	// type aliases
	type metin = string
	type rakam = int8

	var yesim metin = "guzel"
	fmt.Println(yesim)

	// byte versus rune!
	type byte = uint8 // 0..255 // ASCII
	type rune = int32 // 0..COK!

	var a_harfi_1 byte = 65
	var a_harfi_2 byte = 'A'

	fmt.Println(a_harfi_1 == a_harfi_2)
}
