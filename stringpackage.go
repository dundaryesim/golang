package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Compare("umut", "umut"))
	fmt.Println(strings.Contains("kalemlik", "kalem"))
	fmt.Println(strings.ContainsAny("kalemlik", "kalem"))
	fmt.Println(strings.Count("Abacaisa", "a")) // Case sensitive
	fmt.Println(strings.EqualFold("YESIM", "yesim"))
	fmt.Println(strings.Fields("Muazzez Abacı")) //[]
	fmt.Println(strings.HasPrefix("exit", "exw"))
	fmt.Println(strings.HasSuffix("pikaçu", "kaçu"))
	fmt.Println(strings.Index("serhat", "r")) // Count from "0"
	fmt.Println(strings.Repeat("Karıma aşığım \n", 5))
	fmt.Println(strings.Replace("Go Go Go", "Go", "Golang", 2))
	fmt.Println(strings.ReplaceAll("Go Go Go", "Go", "Golang"))
	fmt.Println(strings.Split("a,b,c,d,e", ","))
	fmt.Println(strings.Title("merhaba dunya"))
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "pijamalı hasta yağız şoföre çabucak güvendi"))
	fmt.Println(strings.ToLower("HELLO woRld"))
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Pijamalı haSTa yaĞız şoföre çabucak güvendi"))
	fmt.Println(strings.ToUpper("merhaba dunya"))
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "Pijamalı haSTa yaĞız şoföre çabucak güvendi"))
	fmt.Println(strings.Trim("!!!> Merhaba >!!!", "!>"))

}
