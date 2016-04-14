package fsm

import (
	. "../defines"
	. "../driver"
	. "../network"
	. "../queue"
	"fmt"
	"time"
)

func Fsm(floorCh chan int, buttonPressCh chan ButtonPress, networkCh chan MSG) {

	var doorCloseCh <-chan time.Time

	var Msg = MSG{}

	Msg.State = INIT
	Msg.Floor = -1
	Msg.Dir = DIR_STOP
	Msg.IP = GetLocalIP()

	if ElevGetFloorSensorSignal() == -1 {
		ElevSetMotorDir(DIR_DOWN)
		Msg.Dir = DIR_DOWN
		Msg.State = MOVING
	} else {
		Msg.State = IDLE
	}

	fmt.Println("Fsm started")

	for {

		select {

		case n := <-networkCh:
			switch n {
			case NewOrder:

				break
			case CompletedOrder:

				break
			case ElevatorState:

				break
			}

		case f := <-floorCh:
			Msg.Floor = f
			ElevSetFloorIndicator(f)
			fmt.Println("Event: Arrived at floor", f)
			if QueueShouldStop(&Msg) {
				fmt.Println("Stopping")
				ElevSetMotorDir(DIR_STOP)
				ElevSetDoorOpenLamp(LIGHT_ON)

				Msg.Floor = f

				Msg.State = DOORSOPEN
				doorCloseCh = time.After(3 * time.Second)
			}
			break

		case <-doorCloseCh:
			fmt.Println("Event: Closing door")
			ElevSetDoorOpenLamp(LIGHT_OFF)
			for i := 0; i > 3; i++ {
				ElevSetButtonLamp(i, Msg.Floor, LIGHT_OFF)
			}

			Msg.Dir = QueueChooseDirection(&Msg)
			ElevSetMotorDir(Msg.Dir)
			QueueDeleteCompleted(&Msg)

			if Msg.Dir == DIR_STOP {
				Msg.State = IDLE
			} else {
				Msg.State = MOVING
			}
			break

		case b := <-buttonPressCh:
			fmt.Println("Event: button press: {Floor:", b.Floor, ", Button:", b.Button, "}")
			PrintMsg(&Msg)
			switch Msg.State {
			case MOVING:
				QueueAddOrder(&Msg, b.Floor, b.Button)
				break

			case IDLE:
				if Msg.Floor == b.Floor {
					fmt.Println("Idle -> DoorsOpen")
					ElevSetDoorOpenLamp(LIGHT_ON)
					doorCloseCh = time.After(3 * time.Second)
					Msg.State = DOORSOPEN
				} else {
					fmt.Println("Idle -> Move")
					QueueAddOrder(&Msg, b.Floor, b.Button)
					Msg.Dir = QueueChooseDirection(&Msg)
					ElevSetMotorDir(Msg.Dir)
					Msg.State = MOVING
				}
				break

			case INIT:
				QueueAddOrder(&Msg, b.Floor, b.Button)
				break

			case DOORSOPEN:
				if Msg.Floor == b.Floor {
					fmt.Println("DoorsOpen -> DoorsOpen")
					doorCloseCh = time.After(3 * time.Second)
				} else {
					QueueAddOrder(&Msg, b.Floor, b.Button)
				}
			}
		}
	}
}
