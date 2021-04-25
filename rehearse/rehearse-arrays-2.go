package main

import (
	"fmt"
	"strings"
)

func main() {
	insanlar := [3]string{"serhat", "yesim", "umut"}

	// bir array'i fonksiyona arguman olarak verdiginde, Go onun bir
	// kopyasini olusturur! dolayisiyla array degismez!
	buyukHarf(insanlar)

	fmt.Println(insanlar)
}

// [3]string is the type! uzunlukta turun parcasi! [2]string veremezsin!
func buyukHarf(arr [3]string) {
	for sira, item := range arr {
		arr[sira] = strings.ToUpper(item)
	}
}
