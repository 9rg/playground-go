package carMain

import "fmt"

type Car struct {
	name    string
	owner   string
	fuel    int
	fueleco int
}

func NewCar(name, owner string, fuel, fueleco int) *Car {
	c := new(Car)
	c.name = name
	c.owner = owner
	c.fuel = fuel
	c.fueleco = fueleco
	return c
}

func (c *Car) Run(dis int) {
	if useFuel := dis / c.fueleco; useFuel > c.fuel {
		fmt.Printf("Not enough fuel. remaining fuel:%d\n", c.fuel)
	} else {
		fmt.Printf("%s runs %dkm.\n", c.name, dis)
		c.fuel -= useFuel
	}
}

func (c Car) PrintOwner() {
	fmt.Printf("%s's owner is %s.\n", c.name, c.owner)
}

func (c Car) PrintFuel() {
	fmt.Printf("remaining fuel:%dL.\n", c.fuel)
}
