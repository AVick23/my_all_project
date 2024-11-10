package main

import "fmt"

type Vehicles interface {
	Speed() string
	FuelType() string
}

type Car struct{}

type Bicycle struct{}

type Airplane struct{}

func (a Car) Speed() string {
	return "85 km/h"
}

func (a Car) FuelType() string {
	return "Oil - 95"
}

func (b Bicycle) Speed() string {
	return "20 km/h"
}

func (b Bicycle) FuelType() string {
	return "Food"
}

func (c Airplane) Speed() string {
	return "400 km/h"
}

func (c Airplane) FuelType() string {
	return "Diesel"
}

func PrintInfo(d Vehicles) {
	fmt.Printf("Speed: %s, FuelType: %s \n", d.Speed(), d.FuelType())

}

func main() {
	var (
		Car      = Car{}
		Bicycle  = Bicycle{}
		Airplane = Airplane{}
	)

	PrintInfo(Car)
	PrintInfo(Bicycle)
	PrintInfo(Airplane)
}
