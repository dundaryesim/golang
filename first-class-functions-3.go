package main

import "fmt"

func main() {
	squareInfo(length, area, 10, 20)

	fmt.Println(squareInfoReturn(length, area, 30, 30))
}

func squareInfo(lenghtFunc func(int, int) int, areaFunc func(int, int) int, height int, width int) {
	fmt.Println("Uzunluk: ", lenghtFunc(height, width), "alan: ", areaFunc(height, width))
}

func squareInfoReturn(lenghtFunc func(int, int) int, areaFunc func(int, int) int, height int, width int) (int, int) {
	return lenghtFunc(height, width), areaFunc(height, width)
}

func length(x, y int) int {
	return 2 * (x + y)
}

func area(x, y int) int {
	return x * y
}
