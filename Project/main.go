package main

import (
	"../Project/driver"
	"../Project/fsm"
	"../Project/queue"
	"time"
	"fmt"
)

const N_FLOORS = 4

func main() {
	//Initialize hardware
	if driver.ElevInit() == 0 {
		fmt.Println("Unable to initialize elevator hardware!\n")
		return
	}



	fsm.FsmStartUp()

	for{

		var prevButtons [4][3]int
		for i := 0; i < 4; i++{
			for btn := 0; btn < 3; btn++{
				button bool = driver.ElevgetButtonSignal(btn,i)
				if prevButtons[i][btn] != button && button == true{
					fsm.FsmEvButtonPressed(btn,i)
				}
				prevButtons[i][btn] = button

			}

		}
		queue.QueueSetLights()
		fsm.FsmSetIndicator()

		var prevOrderExist bool
		o bool = queue.QueueOrderExist()
		if o != prevOrderExist{
			if o == true{
				fsm.FsmEvOrderExist()
			}
		}
		prevOrderExist = o

		

		var prevCorrectFloorReached int = -1
		f int = driver.ElevGetFloorSensorSignal()
		if f!= prevCorrectFloorReached && f != -1{
			fsm.FsmEvCorrectFloorReached(f)
		}
		prevCorrectFloorReached = f


		var prevTimeOut int
		t int = 3
		if t != prevTimeOut{
			if t == 3{
				fsm.FsmEvTimeOut()
			}
		}
		prevTimeOut = t


		var prevStopButton int
		s int = fsm.FsmStopButtonIsPressed()
		if s != prevStopButton{
			if s == true{
				fsm.FsmEvStopButtonPressed()
			}
			else{
				fsm.FsmEvStopButtonReleased()
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
