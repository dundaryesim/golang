package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	terrestrial := planets[0:4]
	worlds := append(terrestrial, "Ceres")
	fmt.Println(planets)
	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")
	fmt.Println(planets)
	fmt.Println(worlds)
	fmt.Println(terrestrial)
}
