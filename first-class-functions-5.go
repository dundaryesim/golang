package main

import "fmt"

func main() {
	square := func(x int) int {
		return x * x
	}

	fmt.Println(square(10))

	func(x int) {
		fmt.Println(x / 2)
	}(10)
}
