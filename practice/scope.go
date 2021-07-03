package main

import "fmt"

func main() {
	// var size int // 0
	// var name string    // ""
	// var weight float64 // 0
	// var retired bool // false

	// narrow scope!
	if size := 10; size > 15 {
		fmt.Println("Buyuk")
	} else {
		fmt.Println("Kucuk")
	}
}
