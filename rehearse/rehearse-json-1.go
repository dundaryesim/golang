package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type city struct {
		Name       string
		Population int
	}

	berlin := city{Name: "Berlin", Population: 3769000}

	// marshal returns the JSON encoding of berlin.
	bytes, err := json.Marshal(berlin)

	if err != nil {
		return
	}

	fmt.Println(string(bytes))
}
