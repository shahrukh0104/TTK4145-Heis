package fsm

import(
	."../driver"
	."../defines"
	"time"
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

func FsmEvOrderExist() {
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
			time.Sleep(30 *time.Millisecond)
			}
		break

	case IDLE:
		ElevSetMotorDir(DIR_STOP)
		direction = DIR_STOP
		ElevSetDoorOpenLamp(LIGHT_ON)
		
		thisState = DOORSOPEN
		time.Sleep(30 *time.Millisecond)
		break

	case INIT:
		ElevSetDoorOpenLamp(LIGHT_ON)
		
		thisState = DOORSOPEN
		time.Sleep(30 *time.Millisecond)
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
		
		time.Sleep(30 *time.Millisecond)
		ElevSetDoorOpenLamp(LIGHT_OFF)

		for i :=0; i > 3; i++ {
			ElevSetButtonLamp(i, currentFloor, LIGHT_OFF)
		}

		direction = QueueChooseDirection(currentFloor, direction)
		ElevSetMotorDir(direction)
		QueueDeleteCompleted(currentFloor, direction)
		
		
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

func FsmEvStopButtonPressed(){
	switch thisState{
	case IDLE:
		
		ElevSetMotorDir(DIR_STOP)
		direction = DIR_STOP
		ElevSetStopLamp(LIGHT_ON)
		ElevSetDoorOpenLamp(LIGHT_ON)
		QueueDeleteAllOrders()
	
		thisState = STOP
		break
		
	case MOVING:
		
		ElevSetMotorDir(DIR_STOP)
		direction = DIR_STOP
		ElevSetStopLamp(LIGHT_ON)
		QueueDeleteAllOrders()
	
		thisState = STOP
		break
		
	case DOORSOPEN:
		
		ElevSetMotorDir(DIR_STOP)
		direction = DIR_STOP
		ElevSetStopLamp(LIGHT_ON)
		ElevSetDoorOpenLamp(LIGHT_ON)
		QueueDeleteAllOrders()
		
		thisState = STOP
		break
		
	case INIT:
		ElevSetStopLamp(LIGHT_ON)
		if currentFloor == ElevGetFloorSensorSignal(){
			ElevSetDoorOpenLamp(LIGHT_ON)
		}else{
			ElevSetDoorOpenLamp(LIGHT_OFF)
		}

		thisState = STOP
		break
	
	case STOP:
	
		ElevSetStopLamp(LIGHT_ON)
		ElevSetDoorOpenLamp(LIGHT_OFF)
		
		thisState = STOP
		break
	}
}

func FsmEvStopButtonReleased(){
	switch thisState{
	case STOP:

		ElevSetStopLamp(LIGHT_OFF)
		if currentFloor == ElevGetFloorSensorSignal(){
			ElevSetDoorOpenLamp(LIGHT_ON)
		}
		time.Sleep(30 *time.Millisecond)
		ElevSetDoorOpenLamp(LIGHT_OFF)
			
		/*if currentFloor == driver.ElevGetFloorSensorSignal(){
			for{
				if timer.TimerIsTimeOut == true{
					break
				}
				driver.ElevSetDoorOpenLamp(ON)
			}
		}
                    
		driver.ElevSetDoorOpenLamp(OFF)
		timer.TimerStop()
	*/
		thisState = INIT
		break
		
	case MOVING:
		break
	
	case DOORSOPEN:
		break
	
	case INIT:
		break	
	case IDLE:
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

func FsmStopButtonisPressed() bool {
	if ElevGetStopSignal() == 1{
		return true
	}else{
		return false
	}
}