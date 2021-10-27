package main

import "fmt"

func main() {

	num := 4

	totalnum := num

	for i := 1; i < num; i++ {
		totalnum = totalnum * (num - i)
	}
	fmt.Println(totalnum)

}
