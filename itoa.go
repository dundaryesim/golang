package main

import (
	"fmt"
	"strconv"
)

//int'ı string'e dönüştürme:
func main() {
	countdown1 := 10
	str1 := "Launch in T minus " + strconv.Itoa(countdown1) + " seconds."
	fmt.Println(str1)

	countdown, err := strconv.Atoi("10")
	if err != nil {
		// oh no, something went wrong
	}
	fmt.Println(countdown)
	// diğer bir yol:
	countdown2 := 9
	str2 := fmt.Sprintf("Launch in T minus %v seconds.", countdown2)
	fmt.Println(str2)

}
