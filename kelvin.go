package main

import "fmt"

// kelvinToCelsius converts oK to oC
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
func main() {
	kelvin := 294.0
	kelvin = 233
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "o K is ", celsius, "o C")
}
