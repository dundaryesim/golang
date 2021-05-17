package main

import "fmt"

func main() {
	cities := []string{"Ankara", "Adana", "Bursa", "Balikesir", "Samsun"}
	groupOfSlice := make(map[string][]string)

	for _, v := range cities {
		firstCharacter := string(v[0])

		groupOfSlice[firstCharacter] = append(groupOfSlice[firstCharacter], v)
	}

	for k, v := range groupOfSlice {
		fmt.Println(k, "===>", v)
	}
}
