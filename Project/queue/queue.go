package queue

import (
	. "../defines"
	. "../driver"
)

func QueueOrderExists(States *States) bool {
	for i := 0; i < N_FLOORS; i++ {
		if States.OrderInside[i] == 1 || States.OrderUp[i] == 1 || States.OrderDown[i] == 1 {
			return true
		}
	}
	return false
}

func QueueSetLights(States *States) {
	for i := 0; i < N_FLOORS; i++ {
		ElevSetButtonLamp(BUTTON_COMMAND, i, States.OrderInside[i])
		if i != 0 {
			ElevSetButtonLamp(BUTTON_CALL_DOWN, i, States.OrderDown[i])
		}
		if i != 3 {
			ElevSetButtonLamp(BUTTON_CALL_UP, i, States.OrderUp[i])
		}
	}
}

func QueueAddOrder(States *States, buttonFloor int, buttonTypePressed int) {

	if buttonTypePressed == BUTTON_CALL_UP {
		States.OrderUp[buttonFloor] = 1
	}
	if buttonTypePressed == BUTTON_CALL_DOWN {
		States.OrderDown[buttonFloor] = 1
	}
	if buttonTypePressed == BUTTON_COMMAND {
		States.OrderInside[buttonFloor] = 1
	}
	QueueSetLights(States)
}

func QueueOrdersAbove(States *States) bool {
	for f := States.Floor + 1; f < 4; f++ {
		if States.OrderInside[f] == 1 || States.OrderUp[f] == 1 || States.OrderDown[f] == 1 {
			return true
		}
	}
	return false
}

func QueueOrdersBelow(States *States) bool {
	for f := 0; f < States.Floor; f++ {
		if States.OrderInside[f] == 1 || States.OrderUp[f] == 1 || States.OrderDown[f] == 1 {
			return true
		}
	}
	return false
}

func QueueChooseDirection(States *States) int {
	if States.Dir == DIR_UP {
		if QueueOrdersAbove(States) {
			return DIR_UP
		} else if QueueOrdersBelow(States) {
			return DIR_DOWN
		} else {
			return DIR_STOP
		}
	}

	if States.Dir == DIR_DOWN {
		if QueueOrdersBelow(States) {
			return DIR_DOWN
		} else if QueueOrdersAbove(States) {
			return DIR_UP
		} else {
			return DIR_STOP
		}
	}

	if States.Dir == DIR_STOP {
		if QueueOrdersAbove(States) {
			return DIR_UP
		} else if QueueOrdersBelow(States) {
			return DIR_DOWN
		} else {
			return DIR_STOP
		}
	}
	return DIR_STOP
}

func QueueShouldStop(States *States) bool {
	if States.Dir == DIR_DOWN {
		if States.OrderDown[States.Floor] == 1 || States.OrderInside[States.Floor] == 1 || QueueOrdersBelow(States) == false || States.Floor == 0 {
			return true
		}
	} else {
		if States.OrderUp[States.Floor] == 1 || States.OrderInside[States.Floor] == 1 || QueueOrdersAbove(States) == false || States.Floor == 3 {
			return true
		}
	}
	return false
}

func QueueDeleteAllOrders(States *States) {
	for i := 0; i < N_FLOORS; i++ {
		States.OrderUp[i] = 0
		States.OrderDown[i] = 0
		States.OrderInside[i] = 0
	}
	QueueSetLights(States)
}

func QueueDeleteCompleted(States *States) {

	States.OrderInside[States.Floor] = 0
	States.OrderUp[States.Floor] = 0
	States.OrderDown[States.Floor] = 0

	QueueSetLights(States)
}
