package main

import "fmt"

func main() {
	names := []string{
		"yesim",
		"serhat",
		"umut", // trailing comma
	}

	fmt.Println("length ==> ", len(names), "capacity ==>", cap(names))

	names = append(names, "ozgur")

	fmt.Println("length ==> ", len(names), "capacity ==>", cap(names))

	names = append(names, "ulvi", "selma", "cakil")
}
