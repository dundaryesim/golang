package main

import "fmt"

func main() {
	names1 := []string{
		"yesim",
		"serhat",
	}

	fmt.Println("length ==> ", len(names1), "capacity ==>", cap(names1))

	names2 := append(names1, "umut")

	fmt.Println("length ==> ", len(names2), "capacity ==>", cap(names2))

	names2 = append(names2, "ulvi")

	// names1 and names2 have totally different underlying arrays now
	// because names1 didn't have enough capacity, and appending something
	// to this slice, caused creation of a totally independent copy.
	fmt.Println(names1)
}
