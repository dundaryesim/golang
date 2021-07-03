package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// ascii = byte = uint8 (255 chars)
	// utf8 = rune = int32 (2147483647 chars)

	question := "¿Cómo estás?"

	fmt.Println(len(question), "bytes") // 15 bytes

	fmt.Println(utf8.RuneCountInString(question), "runes") // 12 runes

	c, size := utf8.DecodeRuneInString(question)

	fmt.Printf("First rune: %c %v bytes\n", c, size)
}
