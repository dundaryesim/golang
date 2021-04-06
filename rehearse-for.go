package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		fmt.Println(i * 2)
	}

	for x := 0; ; {
		x++
		fmt.Println(x)

		if x == 30 {
			break
		}
	}
}
