package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal     uint16 // min 0 max 65535
	brakePedal   uint16
	steeringWeel int16 // -32k ---> +32k
	topSpeedKmh  float64
}

func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}

// value receiver method
func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / kmhMultiple)
}

// Pointer method
func (c *car) newTopSpeed(newspeed float64) {
	c.topSpeedKmh = newspeed
}

func newerTopSpeed(c car, speed float64) car {
	c.topSpeedKmh = speed
	return c
}

func main() {
	aCar := car{gasPedal: 65000, brakePedal: 0, steeringWeel: 12561, topSpeedKmh: 225.0}
	// bCar := car{22341, 0, 12562, 225.0}
	fmt.Println(aCar.gasPedal)
	// fmt.Println(bCar.gasPedal)

	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	// fmt.Println(bCar.mph())
	// fmt.Println(aCar.topSpeedKmh)
	aCar.newTopSpeed(500.0)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())

}
