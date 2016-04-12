package queue

import(
	. "../driver"
	. "../defines"
)


func QueueOrderExists(Msg *MSG) bool { 
	for i := 0; i < N_FLOORS; i++{
		if Msg.OrderInside[i] == 1 || Msg.OrderUp[i] == 1 || Msg.OrderDown[i] == 1{
			return true
		}
	}
	return false
}

func QueueSetLights(Msg *MSG) {
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

func QueueAddOrder(Msg *MSG, buttonFloor int, buttonTypePressed int) {

	if buttonTypePressed == BUTTON_CALL_UP{
		Msg.OrderUp[buttonFloor] = 1
	}
	if buttonTypePressed == BUTTON_CALL_DOWN{
		Msg.OrderDown[buttonFloor] = 1
	}
	if buttonTypePressed == BUTTON_COMMAND{
		Msg.OrderInside[buttonFloor] = 1
	}
	QueueSetLights(Msg)
}

func QueueOrdersAbove(Msg *MSG) bool{
	for f := Msg.Floor + 1;  f < 4; f++ {
		if Msg.OrderInside[f] == 1 || Msg.OrderUp[f] == 1 || Msg.OrderDown[f] == 1{
			return true
		} 
	}
	return false
}

func QueueOrdersBelow(Msg *MSG) bool{
	for f := 0; f < Msg.Floor; f++{
		if Msg.OrderInside[f] == 1  || Msg.OrderUp[f] == 1 || Msg.OrderDown[f] == 1{
			return true
		}
	}
	return false
}

func QueueChooseDirection(Msg *MSG) int {
	if Msg.Dir == DIR_UP{
		if QueueOrdersAbove(Msg){
			return DIR_UP
		}else if QueueOrdersBelow(Msg){
			return DIR_DOWN
		}else{
			return DIR_STOP
		}
	}

	if Msg.Dir == DIR_DOWN{
		if QueueOrdersBelow(Msg){
			return DIR_DOWN
		}else if QueueOrdersAbove(Msg){
			return DIR_UP
		}else{
			return DIR_STOP
		}
	}

	if Msg.Dir == DIR_STOP{
		if QueueOrdersAbove(Msg){
			return DIR_UP
		}else if QueueOrdersBelow(Msg){
			return DIR_DOWN
		}else{
			return DIR_STOP
		}
	}
	return DIR_STOP
}


func QueueShouldStop(Msg *MSG) bool {
	if Msg.Dir == DIR_DOWN {
		if Msg.OrderDown[Msg.Floor] == 1  || Msg.OrderInside[Msg.Floor] == 1 || QueueOrdersBelow(Msg) == false || Msg.Floor == 0{
			return true
		}
	} else {
		if Msg.OrderUp[Msg.Floor] == 1 || Msg.OrderInside[Msg.Floor] == 1 || QueueOrdersAbove(Msg) == false || Msg.Floor == 3{
			return true
		}
	}
	return false
}



func QueueDeleteAllOrders(Msg *MSG) {
	for i :=0 ; i < N_FLOORS; i++ {
		Msg.OrderUp[i] 		= 0
		Msg.OrderDown[i]	= 0
		Msg.OrderInside[i]	= 0 
	}
	QueueSetLights(Msg)
}

func QueueDeleteCompleted(Msg *MSG){
	
	Msg.OrderInside[Msg.Floor]	= 0
	Msg.OrderUp[Msg.Floor]		= 0
	Msg.OrderDown[Msg.Floor]	= 0
	
	QueueSetLights(Msg)	
}

