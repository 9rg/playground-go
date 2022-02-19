package main

import (
	"github.com/9rg/playground-go/properness/internal/app/carMain"
)

func main() {
	myCar := carMain.NewCar("bentz", "taro", 25, 10)
	myCar.PrintOwner()
	myCar.PrintFuel()

	myCar.Run(100)
	myCar.PrintFuel()

	myCar.Run(120)
	myCar.PrintFuel()

	myCar.Run(80)
}
