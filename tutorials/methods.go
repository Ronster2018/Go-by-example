package main

import "fmt"

/* Two Types of methods
Value receivers: Just receive valiues and does calculations on those values
Pointer Receivers: If you want to change the value, you use the pointer *
*/

const usixteebitmax float64 = 65535
const kmh_multiple float64 = 1.60934 // 1 mph = 1.6kmh

type car struct {
	gas_pedal      uint16 // min 0, max 65535
	break_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

// Value Receivers
// method that calculates kmh
// associating a method to a struct
func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteebitmax)
}
func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteebitmax / kmh_multiple)
}

// Pointer Receivers: Modifying thee values in our struct.
func (c *car) new_top_speed(new_speed float64) {
	c.top_speed_kmh = new_speed
}

func main() {
	a_car := car{
		gas_pedal:      65000,
		break_pedal:    0,
		steering_wheel: 12561,
		top_speed_kmh:  225.0,
	}

	fmt.Println("Top Speed: ", a_car.top_speed_kmh)
	fmt.Println("Pressure on Pedal: ", a_car.gas_pedal)
	fmt.Println("KMH: ", a_car.kmh())
	fmt.Println("MPH: ", a_car.mph())
	a_car.new_top_speed(500)
	fmt.Println("Top Speed:", a_car.top_speed_kmh)
	fmt.Println("KMH: ", a_car.kmh())
	fmt.Println("MPH: ", a_car.mph())
}
