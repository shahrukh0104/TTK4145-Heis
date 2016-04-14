package defines

import (
	"fmt"
)

const (
	//QUANTITY
	N_FLOORS    int = 4
	N_ELEVATORS int = 1
	N_BUTTONS   int = 3

	//DIRECTIONS
	DIR_UP   int = 1
	DIR_DOWN int = -1
	DIR_STOP int = 0

	//LIGHTS
	LIGHT_ON  int = 1
	LIGHT_OFF int = 0

	//LAMP CALL
	BUTTON_CALL_UP   int = 0
	BUTTON_CALL_DOWN int = 1
	BUTTON_COMMAND   int = 2

	//STATES
	INIT      int = 0
	IDLE      int = 1
	MOVING    int = 2
	DOORSOPEN int = 3
	STOP      int = 4

	//ELEVATOR TYPES
	ELEVTYPE_COMEDI     int = 0
	ELEVTYPE_SIMULATION int = 1
)

type States struct {
	State       int
	Floor       int
	Dir         int
	OrderUp     [N_FLOORS]int
	OrderDown   [N_FLOORS]int
	OrderInside [N_FLOORS]int
	IP          string
}

func PrintMsg(Msg *MSG) {
	fmt.Println()

	for i := 0; i < N_FLOORS; i++ {
		defer fmt.Println(Msg.OrderDown[i], " ", Msg.OrderUp[i], " ", Msg.OrderInside[i])
	}
	switch Msg.State {
	case INIT:
		fmt.Println("State: INIT")
	case IDLE:
		fmt.Println("State: IDLE")
	case MOVING:
		fmt.Println("State: MOVING")
	case DOORSOPEN:
		fmt.Println("State: DOORSOPEN")
	default:
		fmt.Println("Invalid state: ", Msg.State)
	}
}
