package main

import (
	"fmt"
	"strings"
)

// hyperspace removes the space surrounding worlds
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

planets := []string{"Venus", "Earth", "Mars"}

func main() {
	planets := []string{"    Venus ", "    Earth ", "  Mars"}
	hyperspace(planets)
	fmt.Println(strings.Join(planets, ""))
}
