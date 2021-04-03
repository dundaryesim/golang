package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := big.NewInt(86400)

	fmt.Printf("%T\n", x)

	y := new(big.Int)
	y.SetString("2400000000000000000000000000000000000000000000000", 10)
	fmt.Printf("%T\n", y)
	z := new(big.Int)
	z.Div(y, x)
	fmt.Printf("%T\n değeri %[1]v\n", z)

}
