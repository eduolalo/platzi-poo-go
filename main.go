package main

import (
	"github.com/kalmecak/platzi-poo-go/objects"
)

func main() {

	// Una forma de instanciar un objeto
	var car objects.Car
	car.License = "AMQ123"
	car.Driver = objects.Account{
		Name: "Andrés Herrera",
	}
	car.Passenger = 4
	car.PrintDataCar()

	// Otra forma de instanciar un objeto
	car2 := objects.Car{
		License: "QWE567",
		Driver: objects.Account{
			Name: "Andrea Herrera",
		},
		Passenger: 3,
	}
	car2.PrintDataCar()

	// Instanciar un objeto de tipo UberBlack que hereda de Car, no es necesario
	// implementar un constructor en Golang a menos que se deba ejecutar cierta
	// lógica al crear el objeto.
	var uberBlack objects.UberBlack{
		Car: objects.Car{
			License: "QWE567",
			Driver: objects.Account{
				Name: "Andrea Herrera",	
			},
			Passenger: 3,
		},
		TypeCarAccecpted: map[string]map[string]int{},
		SeatMaterial: []string{},
	}
}
