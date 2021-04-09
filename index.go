//Write a program to print each byte (ASCII character) of "shalom", one character per line.
package main

import "fmt"

func main() {
	character := "shalom"
	for r := 0; r < 6; r++ {
		c := character[r]
		fmt.Printf("%v %[1]c\n", c)
	}

}
