package main

import "fmt"

func main() {
	// preallocating a slice capacity, prevents us from doubling
	// the slice cap everytime!
	names1 := make([]string, 0, 200)
	names1 = append(names1, "yesim", "serhat")

	fmt.Println("length ==> ", len(names1), "capacity ==>", cap(names1))

	names1 = append(names1, "umut")

	fmt.Println("length ==> ", len(names1), "capacity ==>", cap(names1))

	names1 = append(names1, "ulvi", "foo")

	fmt.Println("length ==> ", len(names1), "capacity ==>", cap(names1))

	// names1 and names2 have totally different underlying arrays now
	// because names1 didn't have enough capacity, and appending something
	// to this slice, caused creation of a totally independent copy.
	fmt.Println(names1)
}
