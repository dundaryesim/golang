package main

import "fmt"

func pow(rakam, us int) int {
	if us < 0 {
		return 0
	}

	if us == 0 {
		return 1
	}

	total := 1

	for i := 0; i < us; i++ {
		total = total * rakam
	}

	return total
}

func main() {
	fmt.Println(pow(3, 4))
}
