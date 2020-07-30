package main

import "fmt"

type baseCar struct {
	brand               string
	year                int
	volumeTrunk         int
	isEngineStart       bool
	isWindowOpen        bool
	howFullTrunkPercent float64
}

type smallCar struct {
	baseCar
}

type truck struct {
	baseCar
}

func main() {

	car1 := smallCar{
		baseCar{
			brand:               "toyota",
			year:                2009,
			volumeTrunk:         754,
			isEngineStart:       false,
			isWindowOpen:        false,
			howFullTrunkPercent: 24.5,
		},
	}
	car2 := truck{
		baseCar{
			brand:               "volvo",
			year:                2002,
			volumeTrunk:         8340,
			isEngineStart:       true,
			isWindowOpen:        true,
			howFullTrunkPercent: 72.3,
		},
	}

	fmt.Printf("Status auto in the morning:\nCar1: %v\nCar2: %v\n", car1, car2)

	car1.isEngineStart = !car1.isEngineStart
	car1.isWindowOpen = true
	car1.howFullTrunkPercent = 86.3
	car2.isEngineStart = false
	car2.isWindowOpen = false
	car2.howFullTrunkPercent = 2.7

	fmt.Printf("Status auto in the evening:\nCar1: %v\nCar2: %v\n", car1, car2)
}
