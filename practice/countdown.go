package main

import (
	"fmt"
	"time"
)

func main() {
	var x = 10
	for x > 0 {
		fmt.Println(x)
		time.Sleep(time.Second)
		x--
	}
	fmt.Println("liftooff")
}
