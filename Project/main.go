package main

import (
	."../Project/driver"
	."../Project/defines"
	."../Project/fsm"
	."../Project/queue"
    ."../Project/timer"
	"fmt"
)


func main() {

	fmt.Println("Press STOP button to stop elevator and exit program.\n")

	ElevInit(ELEVTYPE_SIMULATION)
	floorCh := make(chan int)
	buttonPressCh := make(chan ButtonPress)
	go ElevPoller(floorCh, buttonPressCh)
	go Fsm(floorChan, buttonPress)




	for {
		select {
		case floor := <-floorChan:
			fmt.Println("Floor: ", floor)
			switch floor {
			case 0:
				ElevSetMotorDir(DIR_UP)
				IoSetIndicator()
				break
			case N_FLOORS - 1:
				ElevSetMotorDir(DIR_DOWN)
				IoSetIndicator()
				break
			}

		case buttonPress := <-buttonPressChan:
			fmt.Println("Button press: { floor ", buttonPress.Floor, ", button ", buttonPress.Button, "}")

		case stop := <-stopChan:
			fmt.Println("Stop: ", stop)
			if stop {
				ElevSetMotorDir(DIR_STOP)
				return
			}
		}
	}
}







//git add .

//git commit -a -m "Fsm nearly complete"

//git push











