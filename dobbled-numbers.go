package main

import "fmt"

func Maps(x []int) []int {
	for i := 0; i < len(x); i++ {
		x[i] = x[i] * 2
	}
	return x

}
func main() {
	ex := []int{
		1,
		2,
		3,
	}
	fmt.Println(Maps(ex))
}
