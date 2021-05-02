package main

import (
	"fmt"
	"strings"
)

// camel case: not exposed
// first letter is important to see if the func is exposed or not!
// mustafaSerhatDundarFromTurkeySamsun

// camel case: not exposed
// first letter is important to see if the func is exposed or not!
// MustafaSerhatDundarFromTurkeySamsun

// snake case
// mustafa_serhat_dundar_from_turkey_samsun

func upperFirstLetter(names []string) {
	// classic way!
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// post condition can be anything!
	// for i := 0; i < len(names); i = i * 10 {
	// 	fmt.Println(names[i])
	// }

	// short-way with range!
	for i := range names {
		fmt.Println(i, " ==> ", names[i])

		names[i] = strings.Title(names[i])
		fmt.Println(names)
	}

	// for i, v := range names {
	// 	// range returns 2 values. 1) Indice, and the 2) item itself!
	// 	fmt.Println(i, " ==> ", v)

	// 	names[i] = strings.Title(v)

	// iterasyon boyle calisacak, ustteki kod icin!:
	// names[0] = "Yesim"
	// names[1] = "Serhat"
	// names[2] = "Umut"
	// }
}

func main() {
	mySlice := []string{
		"yesim",
		"serhat",
		"umut",
	}

	// Passing a slice to a function still creates a copy, but that copy points
	// to the same underlying array, so slices are modifiable when they are
	// passed to functions!
	upperFirstLetter(mySlice)

	fmt.Println("mySlice ===> ", mySlice)
}
