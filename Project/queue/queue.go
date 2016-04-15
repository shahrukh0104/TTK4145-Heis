package queue

import (
	. "../defines"
	. "../driver"
	. "../network"
)


func QueueOrderExists(e *ElevatorState) bool {
	for i := 0; i < N_FLOORS; i++ {
		if e.OrderInside[i] == 1 || e.OrderUp[i] == 1 || e.OrderDown[i] == 1 {
			return true
		}
	}
	return false
}

func QueueSetLights(e *ElevatorState) {
	for i := 0; i < N_FLOORS; i++ {
		ElevSetButtonLamp(BUTTON_COMMAND, i, e.OrderInside[i])
		if i != 0 {
			ElevSetButtonLamp(BUTTON_CALL_DOWN, i, e.OrderDown[i])
		}
		if i != 3 {
			ElevSetButtonLamp(BUTTON_CALL_UP, i, e.OrderUp[i])
		}
	}
}

func QueueAddOrder(e *ElevatorState, buttonFloor int, buttonTypePressed int) {

	if buttonTypePressed == BUTTON_CALL_UP {
		e.OrderUp[buttonFloor] = 1
	}
	if buttonTypePressed == BUTTON_CALL_DOWN {
		e.OrderDown[buttonFloor] = 1
	}
	if buttonTypePressed == BUTTON_COMMAND {
		e.OrderInside[buttonFloor] = 1
	}
	QueueSetLights(e)
	BackupSavetoFile(e)
}

func QueueOrdersAbove(e *ElevatorState) bool {
	for f := e.Floor + 1; f < 4; f++ {
		if e.OrderInside[f] == 1 || e.OrderUp[f] == 1 || e.OrderDown[f] == 1 {
			return true
		}
	}
	return false
}

func QueueOrdersBelow(e *ElevatorState) bool {
	for f := 0; f < e.Floor; f++ {
		if e.OrderInside[f] == 1 || e.OrderUp[f] == 1 || e.OrderDown[f] == 1 {
			return true
		}
	}
	return false
}

func QueueChooseDirection(e *ElevatorState) int {
	if e.Dir == DIR_UP {
		if QueueOrdersAbove(e) {
			return DIR_UP
		} else if QueueOrdersBelow(e) {
			return DIR_DOWN
		} else {
			return DIR_STOP
		}
	}

	if e.Dir == DIR_DOWN {
		if QueueOrdersBelow(e) {
			return DIR_DOWN
		} else if QueueOrdersAbove(e) {
			return DIR_UP
		} else {
			return DIR_STOP
		}
	}

	if e.Dir == DIR_STOP {
		if QueueOrdersAbove(e) {
			return DIR_UP
		} else if QueueOrdersBelow(e) {
			return DIR_DOWN
		} else {
			return DIR_STOP
		}
	}
	return DIR_STOP
}

func QueueShouldStop(e *ElevatorState) bool {
	if e.Dir == DIR_DOWN {
		if e.OrderDown[e.Floor] == 1 || e.OrderInside[e.Floor] == 1 || QueueOrdersBelow(e) == false || e.Floor == 0 {
			return true
		}
	} else {
		if e.OrderUp[e.Floor] == 1 || e.OrderInside[e.Floor] == 1 || QueueOrdersAbove(e) == false || e.Floor == 3 {
			return true
		}
	}
	return false
}

func QueueDeleteAllOrders(e *ElevatorState) {
	for i := 0; i < N_FLOORS; i++ {
		e.OrderUp[i] = 0
		e.OrderDown[i] = 0
		e.OrderInside[i] = 0
	}
	QueueSetLights(e)
}

func QueueDeleteCompleted(e *ElevatorState) {

	e.OrderInside[e.Floor] = 0
	e.OrderUp[e.Floor] = 0
	e.OrderDown[e.Floor] = 0
	

	QueueSetLights(e)
	BackupSavetoFile(e)
}
