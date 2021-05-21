package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// changing names of attributes for JSON!
	type city struct {
		Name       string `json:"name"`
		Population int    `json:"population"`
		Continent  string `json:"continent"`
	}

	cities := []city{
		{Name: "Berlin", Population: 3769000, Continent: "Europe"},
		{Name: "Istanbul", Population: 15460000, Continent: "Europe"},
		{Name: "Tokio", Population: 9273000, Continent: "Asia"},
	}

	// marshal returns the JSON encoding of cities!
	// MarshalIndent is like Marshal but applies Indent to format the output!
	bytes, err := json.MarshalIndent(cities, "", "  ")

	if err != nil {
		return
	}

	fmt.Println(string(bytes))
}
