package fsm

import(
	."../driver"
	."../defines"
	."../timer"
	."../queue"
	"fmt"
)

var thisState		int 	= Msg.State
var direction 		int		= 0 
var currentFloor 	int		= -1

func FsmStartup() {
	for{
		if ElevGetFloorSensorSignal() != -1{
			break
		}
		ElevSetMotorDir(DIR_DOWN)
		var floor int = ElevGetFloorSensorSignal()
		if floor != 1{
			currentFloor = floor
			ElevSetMotorDir(DIR_STOP)
		}
	}
	thisState = IDLE
	fmt.Println("State:", thisState)
}

func FsmEvOrderExist(){
	switch thisState{
	case INIT:
		direction = QueueChooseDirection(currentFloor, direction)
		ElevSetMotorDir(direction)
		
		thisState = MOVING
		break

	case IDLE:
		direction = QueueChooseDirection(currentFloor, direction)
		ElevSetMotorDir(direction)
		
		thisState = MOVING
		break

	case MOVING:
		break


	case DOORSOPEN:
		ElevSetDoorOpenLamp(LIGHT_OFF)
		direction = QueueChooseDirection(currentFloor, direction)
		ElevSetMotorDir(direction)
		
		thisState = MOVING
		break

	case STOP:
		break
	}
}

func FsmEvCorrectFloorReached(newFloor int) {
	currentFloor = newFloor

	switch thisState{
	case MOVING:
		if QueueShouldStop(currentFloor, direction) == 0{
			ElevSetMotorDir(DIR_STOP)
			ElevSetDoorOpenLamp(LIGHT_ON)
			
			thisState = DOORSOPEN
			Timer()
			}
		break

	case IDLE:
		ElevSetMotorDir(DIR_STOP)
		direction = DIR_STOP
		ElevSetDoorOpenLamp(LIGHT_ON)
		
		thisState = DOORSOPEN
		fmt.Println("State:", thisState)

		Timer()
		break

	case INIT:
		ElevSetDoorOpenLamp(LIGHT_ON)
		
		thisState = DOORSOPEN
		Timer()
		break

	case DOORSOPEN:
		break

	case STOP:
		break
	}
}

func FsmEvButtonPressed(buttonPressed int,floor int){
	switch thisState{
	case MOVING:
		QueueAddOrder(floor, buttonPressed) 
		break
	case IDLE:
		QueueAddOrder(floor, buttonPressed) 
		break
	case INIT: 
		QueueAddOrder(floor, buttonPressed) 
		break
	case DOORSOPEN:
		QueueAddOrder(floor, buttonPressed) 
		break

	case STOP:
		break  
	}

}

func FsmEvTimeOut(){
	switch thisState{
	case DOORSOPEN:
		if Timer() == true{
			ElevSetDoorOpenLamp(LIGHT_OFF)

			for i :=0; i > 3; i++ {
                ElevSetButtonLamp(i, currentFloor, LIGHT_OFF)
                direction = QueueChooseDirection(currentFloor, direction)
                
                ElevSetMotorDir(direction)
                QueueDeleteCompleted(currentFloor, direction)
			}
		}

		if direction == DIR_STOP{
			thisState = IDLE

		}else {
			thisState = MOVING
		}	
		break
			
	case MOVING:
		break
	
	case IDLE:
		break
	
	case STOP:
		break
		
	case INIT:
		break
	}
}


func FsmSetIndicator(){
	for i := 0; i<4; i++{
		if ElevGetFloorSensorSignal() == i {
			ElevSetFloorIndicator(i)
		}
	}
}
