package main

import "fmt"

func main() {
	squareInfo(length, area, 10, 20)
}

func squareInfo(lenghtFunc func(int, int) int, areaFunc func(int, int) int, l, h int) {
	fmt.Println(lenghtFunc(l, h), areaFunc(l, h))
}

func length(x, y int) int {
	return 2 * (x + y)
}

func area(x, y int) int {
	return x * y
}
