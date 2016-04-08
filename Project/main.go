package main

import (
	"../Project/driver"
	"fmt"
)

const N_FLOORS = 4

func main() {
	//Initialize hardware
	if driver.ElevInit() == 0 {
		fmt.Println("Unable to initialize elevator hardware!\n")
		return
	}

	fmt.Println("Press STOP button to stop elevator and exit program.\n")

	driver.ElevSetSpeed(-300)
	floorChan := make(chan int)
	buttonPressChan := make(chan driver.ButtonPress)
	stopChan := make(chan bool)
	go driver.ElevPoller(floorChan, buttonPressChan, stopChan)

	for {
		select {
		case floor := <-floorChan:
			fmt.Println("Floor: ", floor)
			switch floor {
			case 0:
				driver.ElevSetSpeed(300)
				break
			case N_FLOORS - 1:
				driver.ElevSetSpeed(-300)
				break
			}

		case buttonPress := <-buttonPressChan:
			fmt.Println("Button press: { floor ", buttonPress.Floor, ", button ", buttonPress.Button, "}")

		case stop := <-stopChan:
			fmt.Println("Stop: ", stop)
			if stop {
				driver.ElevSetSpeed(0)
				return
			}
		}

		/*
			// Change direction when we reach top/bottom floor
			if driver.Elev_get_floor_sensor_signal() == N_FLOORS-1 {
				fmt.Println("her")
			} else if driver.Elev_get_floor_sensor_signal() == 0 {
			} else if driver.Elev_get_stop_signal() == 1 {
				driver.Elev_set_speed(0)
				break
			}
		*/
	}
	return
}