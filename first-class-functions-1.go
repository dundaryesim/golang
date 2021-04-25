package main

import "fmt"

func oldBalance() int {
	return 10
}

func newBalance() int {
	return 20
}

func main() {
	balance := oldBalance  // fonksiyonu degiskene ata ama calistirma
	fmt.Println(balance()) // simdi calistir ve ekrana yazdir

	// reassign balance variable. you can do this because
	// function signatures are identical
	balance = newBalance
	fmt.Println(balance()) // simdi calistir ve ekrana yazdir

	balanceValue := oldBalance() // fonksiyonu calistir ve dondugu degeri degiskene ata
	fmt.Println(balanceValue)    // ekrana yazdir
}
