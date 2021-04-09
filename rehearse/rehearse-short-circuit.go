package main

import "fmt"

func main() {
	rakam := 10
	// ilk condition saglanirsa digerine bakmaz. kisa devre.
	var x = (rakam < 10) && (rakam > 5)  // pesimist
	var y = (rakam > 5) || (rakam > 100) // optimist

	fmt.Println(x)
	fmt.Println(y)
}
