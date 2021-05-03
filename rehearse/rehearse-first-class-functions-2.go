package main

import "fmt"

func oldBalance() int {
	return 10
}

func newBalance() int {
	return 20
}

func main() {
	// balance := oldBalance  // fonksiyonu degiskene ata ama calistirma
	// fmt.Println(balance()) // simdi calistir ve ekrana yazdir

	// do the same without type inference
	var balance func() int = oldBalance
	fmt.Println(balance())
}
