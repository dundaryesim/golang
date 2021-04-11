package main

import (
	"fmt"
	"math"
)

func main() {
	var bh float64 = 32768
	var h = int16(bh)
	fmt.Println(h)
	fmt.Println(math.MinInt16)

}
