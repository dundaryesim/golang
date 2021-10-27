package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var person map[string]string
	var name string
	var address string

	person = make(map[string]string)

	fmt.Println("Write your name:")
	fmt.Scan(&name)

	fmt.Println("Write your address:")
	fmt.Scan(&address)

	person["name"] = name
	person["address"] = address

	j, err := json.Marshal(person)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(j))
	}
}
