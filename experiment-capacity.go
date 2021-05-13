//Write a program that uses a loop to continuously append an element to a slice. Print out
//the capacity of the slice whenever it changes. Does append always double the capacity
//when the underlying array runs out of room?

package main

import "fmt"

func main() {
	var mySlice []string

	firstCapacity := cap(mySlice)

	for i := 0; i < 10000; i++ {
		mySlice = append(mySlice, "yesim")

		if firstCapacity != cap(mySlice) {
			fmt.Println(cap(mySlice))
			firstCapacity = cap(mySlice)
		}
	}
}
