package main

import (
	"fmt"
)

func main() {
	// the keys of maps can be nearly any type, unlike arrays and slices
	samsun := map[string]int{
		"plaka":     55,
		"nufus":     2000000,
		"yuzolcumu": 30000,
	}

	fmt.Println("Samsun'un plakasi: ", samsun["plaka"], "nufusu: ", samsun["nufus"])

	numbersInEnglish := map[int]string{
		0:  "Zero",
		1:  "One",
		2:  "Two",
		90: "Ninety",
	}

	anket := map[string]bool{
		"evlimi?":      false,
		"vatandasmi?":  true,
		"calisiyormu?": true,
	}

	fmt.Println(anket["calisiyormu?"])

	// bunun value'lari slice olsa daha iyi!!!
	cevapAnahtari := map[bool]string{
		true:  "1, 2, 3, 5, 9",
		false: "4, 6, 7, 8",
	}

	fmt.Println(cevapAnahtari[true])

	// add a key to map
	samsun["ilceleri"] = 100
	fmt.Println(samsun["ilceleri"])

	// maps can be iterated with range
	for k, v := range numbersInEnglish {
		fmt.Println(k, " ====> ", v)
	}

	// what if a key doesn't exist in a map and we call it?
	// it returns the zero value
	fmt.Println(numbersInEnglish[40])

	// check if a key exists in map
	rakim, ok := samsun["rakim"]

	if ok {
		fmt.Println("Samsunun rakimi", rakim)
	} else {
		fmt.Println("Bu deger map'te tanimlanmadi!")
	}

	// daha kisa yolu
	if rakim, ok := samsun["rakim"]; ok {
		fmt.Println("Samsunun rakimi", rakim)
	} else {
		fmt.Println("Bu deger map'te tanimlanmadi!")
	}

	// maps aren't copied
	family := map[string]string{
		"father": "Peter",
		"mother": "Lois",
		"kid":    "Stewie",
	}

	familyGuy := family

	familyGuy["dog"] = "Brian"

	fmt.Println(family)
	fmt.Println(familyGuy)

	// maps can also be initialized with make when composite literal isn't prefered
	cartoon := make(map[string]string)
	fmt.Println(cartoon)

	// grouping data with maps and slices
	names := []string{"Serhat", "Sedat", "Semih", "Ahmet", "Anil", "Burak"}
	groups := make(map[string][]string)

	for _, v := range names {
		ilkHarf := string(v[0])

		groups[ilkHarf] = append(groups[ilkHarf], v)
	}

	for k, v := range groups {
		fmt.Println(k, " ===> ", v)
	}
}
