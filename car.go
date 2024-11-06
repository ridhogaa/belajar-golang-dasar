package main

import "fmt"

type Car struct {
	Name, Color string
	Price       int
}

func addCar(price1, price2 int) int {
	return price1 + price2
}

func (receiver Car) PrintCar() {
	fmt.Print(receiver.Name, receiver.Color, receiver.Price)
}

type ICar interface {
	PrintCar()
}
