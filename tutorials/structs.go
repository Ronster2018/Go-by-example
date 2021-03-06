package main

import "fmt"

type car struct {
	gas_pedal      uint16 // min 0, max 65535
	break_pedal    uint16
	steering_wheel int16
	top_speed_mph  float32
}

func main() {
	a_car := car{
		gas_pedal:      22314,
		break_pedal:    0,
		steering_wheel: 12561,
		top_speed_mph:  225.0,
	}

	fmt.Println(a_car.gas_pedal)
}
