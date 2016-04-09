package queue

import(
	. "../driver"
	. "../defines"
)


func QueueOrderExists() bool { 
	for i := 0; i < N_FLOORS; i++{
		if Msg.OrdersInside[i] != 0 || Msg.OrderUp[i] != 0 || Msg.OrderDown[i] != 0{
			return true
		}
	}
	return false
}

func QueueSetLights() {
	for i := 0; i < N_FLOORS;i++{
		ElevSetButtonLamp(BUTTON_COMMAND, i, Msg.OrdersInside[i])
		if i != 0{
			ElevSetButtonLamp(BUTTON_CALL_DOWN, i, Msg.OrderDown[i])
		}
		if i != 3{
			ElevSetButtonLamp(BUTTON_CALL_UP, i, Msg.OrderUp[i])
		}
	}
}

func QueueAddOrder(floor int, buttonTypePressed int) {
	if buttonTypePressed == BUTTON_CALL_UP{
		Msg.OrderUp[floor] = 1
	}
	if buttonTypePressed == BUTTON_CALL_DOWN{
		Msg.OrderDown[floor] = 1
	}
	if buttonTypePressed == BUTTON_COMMAND{
		Msg.OrdersInside[floor] = 1
	}
}

func QueueOrdersAbove(currentFloor int) bool{
	for f := currentFloor+1;  f<4; f++ {
		if Msg.OrdersInside[f] != 0 || Msg.OrderUp[f] != 0 || Msg.OrderDown[f] != 0{
			return true
		} 
	}
	return false
}

func QueueOrdersBelow(currentFloor int) bool{
	for f := 0; f<currentFloor; f++{
		if Msg.OrdersInside[f] != 0  || Msg.OrderUp[f] != 0 || Msg.OrderDown[f] != 0{
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
		return Msg.OrderDown[floor] || Msg.OrderInside[floor] || QueueOrdersBelow(floor)=false || floor = 0
	}
	if prevDir == 1{
		return Msg.OrderUp[floor] || Msg.OrderInside[floor] || QueueOrdersAbove(floor)=false || floor = 3
	}return 1
}




func QueueDeleteAllOrders() {
	for i :=0 ; i<N_FLOORS; i++ {
		Msg.orderUp[i] 		= 0
		Msg.orderDown[i]	= 0
		Msg.orderInside[i]	= 0 
	}
	QueueSetLights()
}

func DeleteCompleted(floor int, prevDirn int){
	
	Msg.orderInside[floor]	= 0
	Msg.orderUp[floor]		= 0
	Msg.orderDown[floor]	= 0
	
	QueueSetLights()	
}

