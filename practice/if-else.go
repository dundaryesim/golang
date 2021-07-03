// Bir sayı 3'e tam bölünebiliyorsa ekrana fizz yaz.
// Bir sayı 5'e tam bölünebiliyorsa ekrana buzz yaz.
// Bir sayı hem 3'e hem 5'e tam bölünebiliyorsa ekrana fizzbuzz yaz.

package main

import "fmt"

func main() {
	var n = 15

	if n%15 == 0 {
		fmt.Println("fizzbuzz")
	} else if n%5 == 0 {
		fmt.Println("buzz")
	} else if n%3 == 0 {
		fmt.Println("fizz")
	} else {
		fmt.Println("none")
	}

}
