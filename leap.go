package main

import "fmt"

func main() {
	var year = 2000
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap == true { // "if leap == true" means "if leap"
		fmt.Println("February is 29 days.")
	} else {
		fmt.Println("February is 28 days.")
	}
}
