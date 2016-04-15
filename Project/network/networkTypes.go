package network

import (
	"encoding/json"
	"fmt"
	. "../defines"
)

var NEWORDER int = 0
var COMPLETEDORDER int = 1
var ELEVATORSTATE int = 2

type NewOrder struct {
	Floor      int
	ButtonType int
	IP         string
}

type CompletedOrder struct {
	Floor      int
	ButtonType int
}

type ElevatorState struct {
	State       int
	Floor       int
	Dir         int
	OrderUp     [N_FLOORS]int
	OrderDown   [N_FLOORS]int
	OrderInside [N_FLOORS]int
	IP          string
}

func NetworkMsgPacker(msgToNetworkCh chan<- NetworkMsg, newOrderCh <-chan NewOrder, CompletedOrderCh <-chan CompletedOrder, elevatorStateCh <-chan ElevatorState) {
	for {
		select {
		case newOrder := <-newOrderCh:
			buf, _ := json.Marshal(newOrder)
			msgToNetworkCh <- NetworkMsg{Msgtype: NEWORDER, Data: buf}

		case completedOrder := <-CompletedOrderCh:
			buf, _ := json.Marshal(completedOrder)
			msgToNetworkCh <- NetworkMsg{Msgtype: COMPLETEDORDER, Data: buf}

		case elevatorState := <-elevatorStateCh:
			buf, _ := json.Marshal(elevatorState)
			msgToNetworkCh <- NetworkMsg{Msgtype: ELEVATORSTATE, Data: buf}
		}
	}

}
func NetworkMsgUnpacker(msgFromNetworkCh <-chan NetworkMsg, newOrderCh chan<- NewOrder, completedOrderCh chan<- CompletedOrder, elevatorStateCh chan<- ElevatorState) {
	for {
		msg := <-msgFromNetworkCh
		switch msg.Msgtype {
		case NEWORDER:
			var newOrder NewOrder
			json.Unmarshal(msg.Data, &newOrder)
			newOrderCh <- newOrder

		case COMPLETEDORDER:
			var completedOrder CompletedOrder
			json.Unmarshal(msg.Data, &completedOrder)
			completedOrderCh <- completedOrder

		case ELEVATORSTATE:
			var elevatorState ElevatorState
			json.Unmarshal(msg.Data, &elevatorState)
			elevatorStateCh <- elevatorState

		}
	}
}

func PrintElevatorState(e *ElevatorState) {
	fmt.Println()

	for i := 0; i < N_FLOORS; i++ {
		defer fmt.Println(e.OrderDown[i], " ", e.OrderUp[i], " ", e.OrderInside[i])
	}
	switch e.State {
	case INIT:
		fmt.Println("State: INIT")
	case IDLE:
		fmt.Println("State: IDLE")
	case MOVING:
		fmt.Println("State: MOVING")
	case DOORSOPEN:
		fmt.Println("State: DOORSOPEN")
	default:
		fmt.Println("Invalid state: ", e.State)
	}
}
