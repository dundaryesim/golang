package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}
	opportunity := location{lat: -1.9462, long: 354.4734}
	inside := location{lat: 4.5, long: 135.9}
	fmt.Println(opportunity, inside)
	// You can write without field name but don't mix up the order of them.
	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit)
	// "v" with "+" print our field name
	curiosity := location{-4.5895, 137.4417}
	fmt.Printf("%v\n", curiosity)
	fmt.Printf("%+v\n", curiosity)
}
