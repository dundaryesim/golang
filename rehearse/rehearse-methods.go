package main

import "fmt"

// method
type animal string

func (x animal) sound(s string) string {
	return fmt.Sprintf(s)
}

func main() {
	// Type alias!
	type human = string

	var yesim human = "Yesim"
	var lastName = "Kuran"

	// you can do this because human is still a string!
	fmt.Println(yesim + lastName)

	// custom types
	var dog animal = "Cakil"
	var owner string = "Yesim"

	fmt.Println(dog)

	// you can't do this because these are different types now!
	fmt.Println(string(dog) + " " + owner)

	// call the method
	fmt.Println(dog.sound("Hav hav hav!"))
}
