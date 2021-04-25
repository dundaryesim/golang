package main

import "fmt"

func main() {
	var x [3][3]int

	fmt.Println(x)

	x[0][0] = 10
	x[2][0] = 40

	fmt.Println(x)
}
