package main

import "fmt"

// forwarding
type car struct {
	year int
	color
	body
	engine
}

type body struct {
	bodyType string // sedan, coupe, SUV
	doors    int    // 2-3, 4-5, 6-7
}

type engine struct {
	year     int // 2017
	power    int // 150hp
	capacity int // 1600cc
}

type color string

func (e engine) taxRate() float64 {
	return float64((e.power + e.capacity) / 2)
}

func main() {
	engine := engine{power: 150, capacity: 1600, year: 2015}
	body := body{bodyType: "Sedan", doors: 5}

	car := car{
		year:   2018,
		color:  "Red",
		body:   body,
		engine: engine,
	}

	fmt.Println(car.engine.power)
	fmt.Println(engine.taxRate())
	fmt.Println(car.taxRate())
	fmt.Println(car.year)
}
