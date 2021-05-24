package main

import "fmt"

type stock struct {
	name     string
	price    float64
	exchange string
}

type average struct {
	fifty      float64
	twohundred float64
}

func (s stock) distance(a average) (float64, float64) {
	var fiftyDiff, twoHundredDiff float64

	fiftyDiff = s.price - a.fifty
	twoHundredDiff = s.price - a.twohundred

	return fiftyDiff, twoHundredDiff
}

func main() {
	s := stock{name: "Apple", price: 250.45}
	a := average{fifty: 300.30, twohundred: 50.20}

	fmt.Println(s.distance(a))
}
