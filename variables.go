package main

import "fmt"

func main() {
	var (
		isim = "Yesim"
		yas  = 18
		kilo = 51
		//	boy   = 1.62
		//	issiz = true
		endex  = (yas + kilo) / 2
		mature = yas >= 18
	)
	fmt.Printf("Adı: %v, Yası: %v, Endexi: %v \n", isim, yas, endex)
	fmt.Printf("yesim resit mi? %v\n", mature)
	fmt.Printf("Yesim 51 kilo mu? %v\n", kilo != 51)

	fmt.Printf("Yeşim şişman mı? %v \n", 60 < kilo)
}
