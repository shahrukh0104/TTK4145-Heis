package queue

import(
	"fmt"
	"timer"
	"../Project/driver"
	"defines"
)

func QueueOrderExists() bool { 
	for i := 0; i < N_FLOORS; i++{
		if(Msg.OrdersInside[i] == true || Msg.OrderUp[i] == true || Msg.OrderDown[i] == true){
			return true
		}
	}
	return false
}

func QueueSetLights() {
	for i := 0; i < N_FLOORS;i++{
		driver.ElevSetButtonLamp(BUTTON_COMMAND, i, OrdersInside[i])
		if i != 0{
			driver.ElevSetButtonLamp(BUTTON_CALL_DOWN, i, OrderDown[i])
		}
		if i != 3{
			driver.ElevSetButtonLamp(BUTTON_CALL_UP, i, OrderUp[i])
		}
	}
}

func QueueAddOrder(floor int, buttonTypePressed int) {
	if buttonTypePressed == BUTTON_CALL_UP{
		OrderUp[floor] = true
	}
	if buttonTypePressed == BUTTON_CALL_DOWN{
		OrderDown[floor] = true
	}
	if buttonTypePressed == BUTTON_COMMAND{
		OrdersInside[floor] = true
	}
}

func QueueOrdersAbove(currentFloor int) bool{
	for f := currentFloor+1;  f<4; f++ {
		if OrdersInside[f] == true || OrderUp[f] == true || OrderDown[f] == true{
			return true
		} 
	}
	return false
}

func QueueOrdersBelow(currentFloor int) bool{
	for f := 0; f<currentFloor; f++{
		if OrdersInside[f] == true || OrderUp[f] == true || OrderDown[f] == true{
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
		return OrderDown[floor] == true || OrderInside[floor] == true || QueueOrdersBelow(floor)==false || floor == 0
	}
	if prevDir == 1{
		return OrderUp[floor]==true || OrderInside[floor]==true || QueueOrdersAbove(floor)==false || floor == 3
	}
	return 1
}




func QueueDeleteAllOrders() {
	for i :=0 ; i<N_FLOORS; i++ {
		orderUp[i] 		= false
		orderDown[i]	= false
		orderInside[i]	= false 
	}
	QueueSetLights()
}

func DeleteCompleted(floor int, prevDirn int){
	
	orderInside[floor]	= false
	orderUp[floor]		= false
	orderDown[floor]	= false
	
	QueueSetLights()	
}

