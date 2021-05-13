package main

import "fmt"

func ekranaBas(ulke string, yas int, isimler ...string) {
	fmt.Println(ulke, yas, isimler)
}

// once zorunlu argumanlar, sonra opsiyoneller!
func ekranaFarkliBas(isimler ...string, ulke string, yas int) {

}

func main() {
	// sadece variadic functionda 0 veya sonsuz arguman verebilirsin
	ekranaBas("turkiye", 10)
}
