package main

import (
	"fmt"
	"math/big"
)

func main() {
	rakam1 := big.NewInt(5)

	// NewInt arguman olarak maximum int64 aliyor
	// ancak bu rakam int64'e sigmiyor!
	// rakam2 := big.NewInt(24000000000000000000)

	rakam2 := new(big.Int)                       // bigInt with zero value
	rakam2.SetString("24000000000000000000", 10) // assign a value

	fmt.Printf("Rakam1'in degeri: %v, turu %[1]T\n", rakam1)
	fmt.Printf("Rakam2'nin degeri: %v, turu %[1]T\n", rakam2)

	// big rakamlarda integer gibi +, -, *, / kullanilamaz!
	toplam := new(big.Int)
	toplam.Add(rakam1, rakam2)

	fark := new(big.Int)
	fark.Sub(rakam2, rakam1)

	carpim := new(big.Int)
	carpim.Mul(rakam1, rakam2)

	bolum := new(big.Int)
	bolum.Div(rakam2, rakam1)

	fmt.Println("Toplam:", toplam)
	fmt.Println("Fark:", fark)
	fmt.Println("Carpim:", carpim)
	fmt.Println("Bolum:", bolum)

}
