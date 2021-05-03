package main

import (
	"fmt"
	"strings"
)

// Custom slice type
type stringlerSlice []string
type integarlarSlice []int

// assigning a "method"! to it
func (slc stringlerSlice) birlestir(sep string) string {
	return strings.Join(slc, sep)
}

func main() {
	isimler := stringlerSlice{
		"yesim",
		"serhat",
		"umut",
	}

	birlesmisIsimler := isimler.birlestir("**")

	fmt.Println(birlesmisIsimler)
}
