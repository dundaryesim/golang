package main

import "fmt"

func main() {
	names := []string{
		"yesim",
		"serhat",
		"ulvi",
		"selma",
		"cakil",
		"umut",
	}

	fmt.Println("length ==> ", len(names), "capacity ==>", cap(names))

	humans := names[0:4:4]

	fmt.Println("length ==> ", len(humans), "capacity ==>", cap(humans))
}
