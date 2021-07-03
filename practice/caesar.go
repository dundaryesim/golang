//What is the result of the expression c = c - 'a' + 'A' if c is a lowercase 'g'?

package main

import "fmt"

func main() {

	c := 'g'
	c = c - 'a' + 'A'
	fmt.Printf("%c %[1]v\n", c)

}
