package main

import (
	. "./defines"
	. "./driver"
	. "./network"
	. "./queue"
	"os/signal"
	"os"
	"fmt"
	"time"
)

func main() {

	ElevInit(ELEVTYPE_COMEDI)
	peerListCh := make(chan []string)
	sendMsgCh := make(chan NetworkMsg)
	recvMsgCh := make(chan NetworkMsg)
	newOrderSendCh := make(chan NewOrder)
	newOrderRecvCh := make(chan NewOrder)
	completedOrderSendCh := make(chan CompletedOrder)
	completedOrderRecvCh := make(chan CompletedOrder)
	elevatorStateSendCh := make(chan ElevatorState)
	elevatorStateRecvCh := make(chan ElevatorState)

	floorCh := make(chan int, 1)
	buttonPressCh := make(chan ButtonPress, 1)

	go UdpSendAlive()
	go UdpRecvAlive(peerListCh)
	go UdpSendMsg(sendMsgCh)
	go UdpRecvMsg(recvMsgCh)
	go NetworkMsgPacker(sendMsgCh, newOrderSendCh, completedOrderSendCh, elevatorStateSendCh)
	go NetworkMsgUnpacker(recvMsgCh, newOrderRecvCh, completedOrderRecvCh, elevatorStateRecvCh)

	go ElevPoller(floorCh, buttonPressCh)
	

	var hallOrders[N_FLOORS][2]string
	states := make( map[string]*ElevatorState )
	var peerList []string
	localIP := GetLocalIP()

	var doorCloseCh <-chan time.Time



	initialFloor := ElevGetFloorSensorSignal()
	for initialFloor == -1 {
		ElevSetMotorDir(DIR_DOWN)
		initialFloor = ElevGetFloorSensorSignal()
		time.Sleep(50*time.Millisecond)
	}

	ElevSetMotorDir(DIR_DOWN)
	states[localIP] = &ElevatorState {
		State:			IDLE,
		Floor:			initialFloor,
		Dir:			DIR_DOWN,
		OrderUp:    	[N_FLOORS]int{},
		OrderDown:  	[N_FLOORS]int{},
		OrderInside:	[N_FLOORS]int{},
		IP:         	localIP,
	}

	BackupLoadFromFile(states[localIP])
	fmt.Println("Restored orders: ", states[localIP].OrderInside)

	go func () {
		var c = make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		<-c
		ElevSetMotorDir(DIR_STOP)
		fmt.Println("Program killed")
		os.Exit(0)
	}()


	onNewOrder := func (floor int, button int){
		fmt.Println("onNewOrder(", floor, ",", button, ")")
		switch states[localIP].State {
		case MOVING:
			QueueAddOrder(states[localIP], floor, button)
			break

		case IDLE:
			if states[localIP].Floor == floor {
				fmt.Println("Idle -> DoorsOpen")
				ElevSetDoorOpenLamp(LIGHT_ON)
				doorCloseCh = time.After(3 * time.Second)
				states[localIP].State = DOORSOPEN
			} else {
				fmt.Println("Idle -> Move")
				QueueAddOrder(states[localIP], floor, button)
				states[localIP].Dir = QueueChooseDirection(states[localIP])
				ElevSetMotorDir(states[localIP].Dir)
				states[localIP].State = MOVING
			}
			break

		case INIT:
			QueueAddOrder(states[localIP], floor, button)
			break

		case DOORSOPEN:
			if states[localIP].Floor == floor {
				fmt.Println("DoorsOpen -> DoorsOpen")
				doorCloseCh = time.After(3 * time.Second)
			} else {
				QueueAddOrder(states[localIP], floor, button)
			}
		}
	}



	for {

		select {
		case newOrder := <-newOrderRecvCh:
			fmt.Println("Event: NewOrder:", newOrder)
			hallOrders[newOrder.Floor][newOrder.ButtonType] = newOrder.IP
			if newOrder.IP == localIP {
				onNewOrder(newOrder.Floor, newOrder.ButtonType)
			} else {
				ElevSetButtonLamp(newOrder.ButtonType, newOrder.Floor, 1)
			}

				
		case completedOrder := <-completedOrderRecvCh:
			fmt.Println("Event: CompletedOrder:", completedOrder)
			hallOrders[completedOrder.Floor][completedOrder.ButtonType] = ""
			ElevSetButtonLamp(completedOrder.ButtonType, completedOrder.Floor, 0)
			break

		case elevatorState := <-elevatorStateRecvCh:
			fmt.Println("Event: ElevatorState:", elevatorState)
			states[elevatorState.IP] = &elevatorState
			break

		case peerList = <- peerListCh:
			fmt.Println("Event: new peer list:", peerList)
			for floor := 0; floor < N_FLOORS; floor++ {
				for button := 0; button < 2; button++ {
					if hallOrders[floor][button] != "" {
						orderIsBeingServiced := false
						for i:= 0; i < len(peerList); i++ {
							if hallOrders[floor][button] == peerList[i] {
								orderIsBeingServiced = true
							}
						}
						if !orderIsBeingServiced {
							onNewOrder(floor, button)
						}
					}
				}
			}

		case f := <-floorCh:
			states[localIP].Floor = f
			ElevSetFloorIndicator(f)
			fmt.Println("Event: Arrived at floor", f)
			if QueueShouldStop(states[localIP]) {
				fmt.Println("Stopping")
				ElevSetMotorDir(DIR_STOP)
				ElevSetDoorOpenLamp(LIGHT_ON)

				states[localIP].Floor = f

				states[localIP].State = DOORSOPEN

				elevatorStateSendCh <- *states[localIP]
			

				doorCloseCh = time.After(3 * time.Second)
			}
			break

		case <-doorCloseCh:
			fmt.Println("Event: Closing door")
			ElevSetDoorOpenLamp(LIGHT_OFF)
			for i := 0; i > 3; i++ {
				ElevSetButtonLamp(i, states[localIP].Floor, LIGHT_OFF)
			}

			states[localIP].Dir = QueueChooseDirection(states[localIP])
			ElevSetMotorDir(states[localIP].Dir)
			QueueDeleteCompleted(states[localIP])
			for button := 0; button < 2; button++ {
				completedOrderSendCh <- CompletedOrder{Floor: states[localIP].Floor, ButtonType: button}
			}

			if states[localIP].Dir == DIR_STOP {
				states[localIP].State = IDLE
				
				elevatorStateSendCh <- *states[localIP]

			} else {
				states[localIP].State = MOVING
				elevatorStateSendCh <- *states[localIP]
			}
			break

		case b := <-buttonPressCh:
			fmt.Println("Event: button press: {Floor:", b.Floor, ", Button:", b.Button, "}")
			PrintElevatorState(states[localIP])
			switch b.Button {
				case BUTTON_COMMAND:
					onNewOrder(b.Floor, b.Button)


					break
				case BUTTON_CALL_DOWN, BUTTON_CALL_UP:
					fmt.Println("hall button pressed ", peerList)
					ServicerIP := CostGlobal(states, peerList, localIP, b.Floor, b.Button)
					newOrderSendCh <- NewOrder{Floor: b.Floor, ButtonType: b.Button, IP: ServicerIP}
					if(ServicerIP == localIP){
						onNewOrder(b.Floor, b.Button)
					}	
			}
		}
	}	
}
		