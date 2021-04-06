package main

import "fmt"

func main() {
	if x := 10; x < 15 {
		fmt.Printf("15'ten kucuk mu? %v\n", true)
	} else if x < 11 {
		fmt.Printf("11'den kucuk mu? %v\n", true)
	} else {
		fmt.Printf("Something else\n")
	}
}
