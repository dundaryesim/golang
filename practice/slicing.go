package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := planets[0:4] // [Mercury Venus Earth Mars]
	gasGiants := planets[4:6]   // [Jupiter Saturn]
	iceGiants := planets[6:8]   // [Uranus Neptune]

	fmt.Println(terrestrial, gasGiants, iceGiants)
	fmt.Println(gasGiants[0]) // Jupiter

	giants := planets[4:8] // [Jupiter Saturn Uranus Neptune]
	gas := giants[0:2]     // [Jupiter Saturn]
	ice := giants[2:4]     // [Uranus Neptune]
	fmt.Println(giants, gas, ice)

	iceGiantsMarkII := iceGiants // [Uranus Neptune]
	iceGiants[1] = "Poseidon"
	fmt.Println(planets)                         // [Mercury Venus Earth Mars Jupiter Saturn Uranus Poseidon]
	fmt.Println(iceGiants, iceGiantsMarkII, ice) // [Uranus Poseidon] [Uranus Poseidon] [Uranus Poseidon]
	// Slice'lar array gibi başka bir değişkene atandığında kopyasını oluşturmaz.
	allPlanets := planets[:]
	fmt.Println(allPlanets)

}
