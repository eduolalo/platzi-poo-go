package objects

import "log"

// Car representa la "clase" de un auto
type Car struct {
	ID        int
	License   string
	Driver    Account
	Passenger int
}

// PrintDataCar imprime los datos del auto
func (c *Car) PrintDataCar() {
	log.Printf("Info: %+v", c)
}
