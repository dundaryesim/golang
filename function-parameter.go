package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vo K\n", k)
		time.Sleep(time.Second)
	}
}
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func main() {
	measureTemperature(3, fakeSensor)
}
//Rewrite the following function signature to use a function type:
//func drawTable(rows int, getRow func(row int) (string, string))
// //type getRowfn func(row int) (string, string)
// func drawTable(rows int, getRow getRowfn)

