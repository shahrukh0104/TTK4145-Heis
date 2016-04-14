package main

import (
	. "./defines"
	. "./driver"
	. "./fsm"
	. "./network"
	"fmt"
)

func main() {

	ElevInit(ELEVTYPE_COMEDI)

	if ElevGetFloorSensorSignal() == -1 {
		ElevSetMotorDir(DIR_DOWN)
		States.Dir = DIR_DOWN
		States.State = MOVING
	} else {
		States.State = IDLE
	}

	peerListCh := make(chan []string)
	sendMsgCh := make(chan States)
	recvMsgCh := make(chan States)

	floorCh := make(chan int, 1)
	buttonPressCh := make(chan ButtonPress, 1)

	go UdpSendAlive()
	go UdpRecvAlive(peerListCh)
	go UdpSendMsg(sendMsgCh)
	go UdpRecvMsg(recvMsgCh)
	go NetworkMsgPacker(sendMsgCh)
	go NetworkMsgUnpacker(recvMsgCh)

	go ElevPoller(floorCh, buttonPressCh)

	var hallOrders [N_FLOORS][2]string // IP of elevator servicing that order, or emptystring
	var states map[string]States
	var peers []string

	var doorCloseCh <-chan time.Time

	States.State = INIT
	States.Floor = -1
	States.Dir = DIR_STOP
	State.IP = GetLocalIP()

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
			States.Floor = f
			ElevSetFloorIndicator(f)
			fmt.Println("Event: Arrived at floor", f)
			if QueueShouldStop(&States) {
				fmt.Println("Stopping")
				ElevSetMotorDir(DIR_STOP)
				ElevSetDoorOpenLamp(LIGHT_ON)

				States.Floor = f

				States.State = DOORSOPEN

				sendMsgCh <- ElevatorState{Floor: f, Dir: DIR_STOP, State: DOORSOPEN}
				UdpSendMsg(sendMsgCh)

				doorCloseCh = time.After(3 * time.Second)
			}
			break

		case <-doorCloseCh:
			fmt.Println("Event: Closing door")
			ElevSetDoorOpenLamp(LIGHT_OFF)
			for i := 0; i > 3; i++ {
				ElevSetButtonLamp(i, States.Floor, LIGHT_OFF)
			}

			States.Dir = QueueChooseDirection(&States)
			ElevSetMotorDir(States.Dir)
			QueueDeleteCompleted(&States)

			if States.Dir == DIR_STOP {
				States.State = IDLE
				
				sendMsgCh <- ElevatorState{Dir: DIR_STOP, OrdersInside[] = nil, State: IDLE}
				UdpSendMsg(sendMsgCh)

			} else {
				States.State = MOVING
				sendMsgCh <- ElevatorState{Dir: DIR_STOP, OrdersInside[] = nil, State: MOVING}
				UdpSendMsg(sendMsgCh)
			}
			break

		case b := <-buttonPressCh:
			fmt.Println("Event: button press: {Floor:", b.Floor, ", Button:", b.Button, "}")
			PrintMsg(&States)
			switch b{
				case: b.Button == BUTTON_COMMAND

				heis States.IP = costfunction()
				
				sendMsgCh <- NewOrder{Floor: b.Floor, ButtonType: BUTTON_COMMAND, IP: heis}
				UdpSendMsg(sendMsgCh)







			}
			switch States.State {
			case MOVING:
				QueueAddOrder(&States, b.Floor, b.Button)
				break

			case IDLE:
				if States.Floor == b.Floor {
					fmt.Println("Idle -> DoorsOpen")
					ElevSetDoorOpenLamp(LIGHT_ON)
					doorCloseCh = time.After(3 * time.Second)
					States.State = DOORSOPEN
				} else {
					fmt.Println("Idle -> Move")
					QueueAddOrder(&States, b.Floor, b.Button)
					States.Dir = QueueChooseDirection(&States)
					ElevSetMotorDir(States.Dir)
					States.State = MOVING
				}
				break

			case INIT:
				QueueAddOrder(&States, b.Floor, b.Button)
				break

			case DOORSOPEN:
				if States.Floor == b.Floor {
					fmt.Println("DoorsOpen -> DoorsOpen")
					doorCloseCh = time.After(3 * time.Second)
				} else {
					QueueAddOrder(&States, b.Floor, b.Button)
				}
			}
		}
	}
}

//git add .

//git commit -a -m "Fsm nearly complete"

//git push
