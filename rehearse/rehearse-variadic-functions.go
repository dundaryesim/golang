package main

import "fmt"

// you can pass 0 to infinite number of arguments to a variadic function
func ekranaYazdir(isimler ...string) {
	fmt.Println(isimler)
}

func ekranaBaskaYazdir(rakam int, yesim bool, isimler ...string) {
	fmt.Println(rakam, isimler, yesim)
}

func main() {
	ekranaYazdir("serhat", "yesim", "umut")
	ekranaYazdir()

	ekranaBaskaYazdir(10, true, "serhat", "yesim")
}
