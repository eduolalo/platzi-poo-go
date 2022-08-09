package main

import (
	"github.com/kalmecak/platzi-poo-go/objects"
)

func main() {

	// Una forma de instanciar un objeto
	var car objects.Car
	car.License = "AMQ123"
	car.Driver = "Andrés Herrera"
	car.Passenger = 4
	car.PrintDataCar()

	// Otra forma de instanciar un objeto
	car2 := objects.Car{
		License:   "QWE567",
		Driver:    "Andrea Herrera",
		Passenger: 3,
	}
	car2.PrintDataCar()
}
