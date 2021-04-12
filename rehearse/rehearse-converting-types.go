package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var yas = 30
	var kilo = 75.9
	var money = 3250001

	// uint8 veya int32 sart! eger string(..) kullanmak istiyorsan!
	var firstLetter uint8 = 65

	var isim = "Serhat"

	// int to uint8
	fmt.Printf("uint8 olarak yas %T ve degeri %[1]v\n", uint8(yas))

	// int to float32
	fmt.Printf("float olarak yas %T ve degeri %[1]v\n", float32(yas))

	// float64 to int
	fmt.Printf("int olarak kilo %T ve degeri %[1]v\n", int(kilo))

	// a big number to uint8 (wraps!)
	fmt.Printf("uint8 olarak money %T ve degeri %[1]v\n", uint8(money))

	// check if a number fits into limits to prevent wrapping!
	if money > math.MaxUint8 || money < 0 {
		fmt.Println("Hata!")
	} else {
		fmt.Println(uint8(money))
	}

	// byte & rune to string
	// you can't convert other int types to string like this!
	var stringFirstLetter = string(firstLetter)
	fmt.Println(stringFirstLetter)

	// int to string [yontem 1]
	stringYas := fmt.Sprintf("%v", yas)
	fmt.Printf("Deger %v, turu %[1]T\n", stringYas)

	// int to string [yontem 2]
	// Atoi => ASCII to integer
	// Itoa => Integer to ASCII
	stringYas2 := strconv.Itoa(yas)
	fmt.Printf("Deger %v, turu %[1]T\n", stringYas2)

	// string to int [yontem 1]
	intYas, err := strconv.Atoi("30")
	fmt.Printf("Deger %v, turu %[1]T ve error %v\n", intYas, err)

	// string to int [hata version]
	intYas1, err := strconv.Atoi(isim)
	fmt.Printf("Deger %v, turu %[1]T ve error %v\n", intYas1, err)

	if err != nil {
		fmt.Println("Bir hata olustu ve string integer'a cevrilemedi!")
	}

	// boolean to string
	var isMarried = true
	var strBool string

	if isMarried == true {
		strBool = "true"
	} else {
		strBool = "false"
	}

	fmt.Printf("Deger %v, turu %[1]T\n", strBool)

	// string to boolean
	var strBool2 = "false"

	switch strBool2 {
	case "true":
		isMarried = true
	case "false":
		isMarried = false
	}

	fmt.Printf("Deger %v, turu %[1]T\n", isMarried)
}
