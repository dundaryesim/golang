// Write a celsiusToKelvin function that uses the celsius and kelvin types
//defined in listing 13.4. Use it to convert 127 C, the surface temperate of the sunlit moon, to
//degrees Kelvin.
package main

import "fmt"

type celsius float64
type kelvin float64

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var c celsius = 127
	k := celsiusToKelvin(c)
	fmt.Print(c, "o C is ", k, "o K")

}
