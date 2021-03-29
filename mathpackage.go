package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("%v'nin mutlak degeri %v.\n", -10.2, math.Abs(-10.2))
	fmt.Printf("%v'den sonra gelen ilk tamsayı %v'dir.\n", 10.1, math.Ceil(10.1))
	fmt.Printf("9.6 değerinden önce gelen ilk tamsayı %v'dir \n", math.Floor(9.6))
	fmt.Printf("Verilen iki sayıdan büyük olanı %v'dir\n", math.Max(10, -10))
	fmt.Printf("Verilen iki sayıdan küçük olanı %v'dir\n", math.Min(574, 9588495))
	fmt.Printf("%v'nin %v'ye bölümünden kalan %v'dir\n", 11, 3, math.Mod(11, 3))
	fmt.Printf("9 üssü 5 %v eder\n", math.Pow(9, 5))
	fmt.Printf("%v'nin kare kökü %v'dir\n", 6, math.Sqrt(6))

}
