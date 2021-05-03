package main

import "fmt"

func pow(sayı, üs int) int {
	if üs == 0 {
		return 1
	}
	if üs < 0 {
		return 0
	}

	total := 1
	for i := 0; i < üs; i++ {
		total = total * sayı

	}
	return total
}

func main() {

	fmt.Println(pow(2, -1))

}
