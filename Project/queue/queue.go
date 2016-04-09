package queue

import(
	. "../driver"
	. "../defines"
	"fmt"
)


func QueueOrderExists() bool { 
	for i := 0; i < N_FLOORS; i++{
		if Msg.OrderInside[i] != 0 || Msg.OrderUp[i] != 0 || Msg.OrderDown[i] != 0{
			return true
		}
	}
	return false
}

func QueueSetLights() {
	for i := 0; i < N_FLOORS;i++{
		ElevSetButtonLamp(BUTTON_COMMAND, i, Msg.OrderInside[i])
		if i != 0{
			ElevSetButtonLamp(BUTTON_CALL_DOWN, i, Msg.OrderDown[i])
		}
		if i != 3{
			ElevSetButtonLamp(BUTTON_CALL_UP, i, Msg.OrderUp[i])
		}
	}
}

func QueueAddOrder(floor int, buttonTypePressed int) {
	fmt.Println("HERERRER")

	if buttonTypePressed == BUTTON_CALL_UP{
		Msg.OrderUp[floor] = 1
	}
	if buttonTypePressed == BUTTON_CALL_DOWN{
		Msg.OrderDown[floor] = 1
	}
	if buttonTypePressed == BUTTON_COMMAND{
		Msg.OrderInside[floor] = 1
	}
}

func QueueOrdersAbove(currentFloor int) bool{
	for f := currentFloor+1;  f<4; f++ {
		if Msg.OrderInside[f] != 0 || Msg.OrderUp[f] != 0 || Msg.OrderDown[f] != 0{
			return true
		} 
	}
	return false
}

func QueueOrdersBelow(currentFloor int) bool{
	for f := 0; f < currentFloor; f++{
		if Msg.OrderInside[f] != 0  || Msg.OrderUp[f] != 0 || Msg.OrderDown[f] != 0{
			return true
		}
	}
	return false
}

func QueueChooseDirection(currentFloor int, prevDir int) int {
	if prevDir == DIR_UP{
		if QueueOrdersAbove(currentFloor){
			return DIR_UP
		}else if QueueOrdersBelow(currentFloor){
			return DIR_DOWN
		}else{
			return DIR_STOP
		}
	}

	if prevDir == DIR_DOWN{
		if QueueOrdersBelow(currentFloor){
			return DIR_DOWN
		}else if QueueOrdersAbove(currentFloor){
			return DIR_UP
		}else{
			return DIR_STOP
		}
	}

	if prevDir == DIR_STOP{
		if QueueOrdersAbove(currentFloor){
			return DIR_UP
		}else if QueueOrdersBelow(currentFloor){
			return DIR_DOWN
		}else{
			return DIR_STOP
		}
	}
	return DIR_STOP
}


func QueueShouldStop(floor int, prevDir int) int{
	if prevDir == -1{
		if Msg.OrderDown[floor] != 0  || Msg.OrderInside[floor] != 0 || !QueueOrdersBelow(floor) || floor == 0{
			return 0
		}
	}
	if prevDir == 1{
		if Msg.OrderUp[floor] != 0 || Msg.OrderInside[floor] != 0 || !QueueOrdersAbove(floor) || floor == 3{
			return 0
		}
	}
	return 1
}



func QueueDeleteAllOrders() {
	for i :=0 ; i < N_FLOORS; i++ {
		Msg.OrderUp[i] 		= 0
		Msg.OrderDown[i]	= 0
		Msg.OrderInside[i]	= 0 
	}
	QueueSetLights()
}

func QueueDeleteCompleted(floor int, prevDirn int){
	
	Msg.OrderInside[floor]	= 0
	Msg.OrderUp[floor]		= 0
	Msg.OrderDown[floor]	= 0
	
	QueueSetLights()	
}

