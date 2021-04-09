package main

import "fmt"

// package main scope'u
var x = 10

func main() {
	// main function scope'u
	fmt.Println(x)

	var f float64 = 5
	f = f + 2

	fmt.Printf("%05.2f ve turu %T\n", f, f)

	var y = 15

	if 10 < 15 {
		// if'in scope'u
		var z = 25
		fmt.Printf("y'nin degeri %v, z'nin degeri %v\n", y, z)
	}

	switch y {
	case 15:
		// ilk case'in scope'u
		var t = 35
		fmt.Println("15! ve t'nin degeri", t)
	case 20:
		// ikinci case'in scope'u
		fmt.Println("20!")
	}

}

// package main scope'u
