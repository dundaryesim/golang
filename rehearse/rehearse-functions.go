package main

import (
	"fmt"
)

func identity(fname string, lname string, yearOfBirth int) (string, int) {
	fullName := fname + " " + lname
	age := 2021 - yearOfBirth

	return fullName, age
}

func main() {
	fmt.Println(identity("Yesim", "Dundar", 1980))
}
