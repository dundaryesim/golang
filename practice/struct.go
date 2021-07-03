package main

import "fmt"

func main() {
	var curiosity struct {
		lat  float64
		long float64
	}
	curiosity.lat = -45895
	curiosity.long = 137.4427
	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)
}
