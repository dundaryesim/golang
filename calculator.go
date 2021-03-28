package main

import (
	"fmt"
)

func main() {
	fmt.Println(Toplama(10, 5))
	fmt.Println(Toplama(8, 12))
	fmt.Println(Cikarma(5, 1))
	fmt.Println(Carpma(4, 68, 3838, 494994))
	fmt.Println(Bolme(17, 6))
	fmt.Println(DavetliListesi("abu", "dabi", "li"))
}

func Toplama(x int, y int) int {
	return x + y
}

func Cikarma(x int, y int) int {
	return x - y
}

func Carpma(x int, y int, z int, t int) int {
	return x * y * z * t

}

func Bolme(x float64, y float64) float64 {
	return x / y
}

func DavetliListesi(x, y, z string) string {
	return x + " " + y + "  " + z
}
