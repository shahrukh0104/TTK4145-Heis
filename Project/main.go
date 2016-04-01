package main

import (
	"../Project/driver"
	"fmt"
)

const N_FLOORS = 4

func main() {
	//Initialize hardware
	if driver.Elev_init() == 0 {
		fmt.Println("Unable to initialize elevator hardware!\n")
		return
	}

	fmt.Println("Press STOP button to stop elevator and exit program.\n")

	driver.Elev_set_speed(-300)

	for {

		// Change direction when we reach top/bottom floor
		if driver.Elev_get_floor_sensor_signal() == N_FLOORS-1 {
			fmt.Println("her")
			driver.Elev_set_speed(-300)
		} else if driver.Elev_get_floor_sensor_signal() == 0 {
			driver.Elev_set_speed(300)
		} else if driver.Elev_get_stop_signal() == 1 {
			driver.Elev_set_speed(0)
			break
		}
	}
	return
}
