package main

import (
	"fmt"
	"math"
)

func main() {
	rakam1 := 0.1
	rakam2 := 0.2
	rakam3 := 0.3

	fmt.Printf("0.1 + 0.2 ?= 0.3 %v\n", math.Abs(rakam1+rakam2-rakam3) < 0.0001)
}
