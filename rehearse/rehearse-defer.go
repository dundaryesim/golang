package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		//  Deferred calls are executed in last-in-first-out order.
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
