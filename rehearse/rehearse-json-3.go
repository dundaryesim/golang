package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// omitempty removes the empty keys from JSON
	// string returns integers between quotes
	type city struct {
		Name       string   `json:"name,omitempty"`
		Population int      `json:"population,omitempty,string"`
		Continent  []string `json:"continent,omitempty"`
	}

	cities := []city{
		{Name: "Berlin", Population: 3769000, Continent: []string{"Europe"}},
		{Name: "Istanbul", Population: 15460000, Continent: []string{"Europe", "Asia"}},
		{Name: "Tokio", Population: 9273000, Continent: []string{}},
	}

	// Marshal a JSON (bytes)
	bytes, _ := json.MarshalIndent(cities, "", "  ")

	// Convert marshaled JSON to string
	fmt.Println(string(bytes))

	// Check if a JSON is valid
	fmt.Printf("Is this is a valid JSON? %v\n", json.Valid([]byte(bytes)))

	// Unmarshal: Converts marshalled JSON to a new structure
	type yeniSehir struct {
		Isim  string   `json:"name,omitempty"`
		Nufus int      `json:"population,omitempty,string"`
		Kita  []string `json:"continent,omitempty"`
	}

	var yeniSehirler []yeniSehir

	json.Unmarshal(bytes, &yeniSehirler)

	fmt.Printf("%+v", yeniSehirler)
}
