//Write and use a celsiusToFahrenheit temperature conversion function. Hint: the
//formula for converting to Fahrenheit is: (c * 9.0 / 5.0) + 32.0.
//-----------
//Write a kelvinToFahrenheit function and verify that it converts 0 K to approximately –459.67° F

package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0

}

func kelvinToFahrenheit(k float64) float64 {
	return k - 459.67

}

func main() {
	celsius := 40
	fahrenheit := celsiusToFahrenheit(float64(celsius))
	fmt.Printf("%v\u00B0C is %v \u00B0F\n", celsius, fahrenheit)

	kelvin := 500
	fahrenheit = kelvinToFahrenheit(float64(kelvin))
	fmt.Printf("%v\u00B0K is %v \u00B0F\n", kelvin, fahrenheit)

}
