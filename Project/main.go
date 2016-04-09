package main

import (
	."../Project/driver"
	."../Project/fsm"
	."../Project/queue"
	"fmt"
)

const N_FLOORS = 4

func main() {
	fmt.Println("Her")
	if ElevInit() == 0 {
		fmt.Println("Unable to initialize elevator hardware!\n")
		return
	}
	fmt.Println("Her")
	FsmStartup()
	for{
		var prevButtons [4][3]int
		for i := 0; i < 4; i++{
			for btn := 0; btn < 3; btn++{
				var button int = ElevGetButtonSignal(btn,i)
				if prevButtons[i][btn] != button && button == 1{
					FsmEvButtonPressed(btn,i)
					fmt.Println("Her")

				}
				prevButtons[i][btn] = button
			}
		}
		QueueSetLights()
		FsmSetIndicator()

		var prevOrderExist bool
		var o bool = QueueOrderExists()
		if o != prevOrderExist{
			if o == true{
				FsmEvOrderExist()
			}
		}
		prevOrderExist = o

		

		var prevCorrectFloorReached int = -1
		var f int = ElevGetFloorSensorSignal()
		if f!= prevCorrectFloorReached && f != -1{
			FsmEvCorrectFloorReached(f)
		}
		prevCorrectFloorReached = f


		var prevTimeOut int
		var t int = 3
		if t != prevTimeOut{
			if t == 3{
				FsmEvTimeOut()
			}
		}
		prevTimeOut = t


		var prevStopButton bool
		var s bool = FsmStopButtonisPressed()
		if s != prevStopButton{
			if s == true{
				FsmEvStopButtonPressed()
			}else{
				FsmEvStopButtonReleased()
			}
			prevStopButton = s
		}
	}
}





/*

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
			}
	return
	*/
