package main

import "fmt"

func main() {
	// %T tell the type
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)
	//days := 365.2425
	//fmt.Printf("Type %T for %[1]v\n", days)
	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days)

	a := "text"
	fmt.Printf("Type %T for %[1]v\n", a)
	b := 42
	fmt.Printf("Type %T for %[1]v\n", b)
	c := 3.14
	fmt.Printf("Type %T for %[1]v\n", c)
	d := true
	fmt.Printf("Type %T for %[1]v\n", d)

	day := "Monday"
	month := "January"
	years := 2020

	// [1] ornegi:
	fmt.Printf("Type is %T for %[1]v\n", day)

	// [2] ornegi:
	fmt.Printf("Type is %T for %[1]v. Type is %T for %[2]v\n", day, month)

	// [3] ornegi:
	fmt.Printf("Type is %T for %[1]v. Type is %T for %v\n", day, month, years)
}
