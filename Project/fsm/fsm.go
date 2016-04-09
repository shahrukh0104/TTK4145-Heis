package fsm

import(
	"fmt"
	"../driver"
	"time"
	"../queue"
)


thisState 		State 	= Msg.State
direction 		int		= 0 
currentFloor 	int		= -1

func FsmStartup() {
	for{
		if driver.ElevGetFloorSignal() != -1{
			break
		}
		driver.ElevSetMotorDirection(DIR_DOWN)
		floor int = driver.ElevGetFloorSignal()
		if floor != 1{
			currentFloor = floor
			driver.ElevSetMotorDirection(DIR_STOP)
		}
	}
	thisState = IDLE
}

func FsmEvOrderExist() {
	switch thisState{
	case INIT:
		direction = queue.QueueChooseDirection(currentFloor, direction)
		ElevSetMotorDirection(direction)
		
		thisState = MOVING
		break

	case IDLE:
		direction = queue.QueueChooseDirection(currentFloor, direction)
		driver.ElevSetMotorDirection(direction)
		
		thisState = MOVING
		break

	case MOVING:
		break


	case DOORSOPEN
		driver.ElevSetDoorOpenLamp(OFF)
		direction = queue.QueueChooseDirection(currentFloor, direction)
		driver.ElevSetMotorDirection(direction)
		
		thisState = MOVING
		break

	case STOP
		break
	}
}

func FsmEvCorrectFloorReached(newFloor int) {
	currentFloor = newFloor

	switch thisState{
	case MOVING:
		if queue.QueueShouldStop(currentFloor, direction){
			driver.ElevSetMotorDirection(DIR_STOP)
			driver.ElevSetDoorOpenLamp(ON)
			
			thisState = DOORSOPEN
			time.Sleep(3000 *time.Milliseconds)
			}
		break

	case IDLE:
		driver.ElevSetMotorDirection(DIR_STOP)
		direction = DIR_STOP
		driver.ElevSetDoorOpenLamp(ON)
		
		thisState = DOORSOPEN
		time.Sleep(3000 *time.Milliseconds)
		break

	case INIT:
		driver.ElevSetDoorOpenLamp(ON)
		
		thisState = DOORSOPEN
		time.Sleep(3000 *time.Milliseconds)
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
		queue.QueueAddOrder(floor, buttonPressed) 
		break

	case IDLE:
		queue.QueueAddOrder(floor, buttonPressed) 
		break

	case INIT: 
		queue.QueueAddOrder(floor, buttonPressed) 
		break

	case DOORSOPEN:
		queue.QueueAddOrder(floor, buttonPressed) 
		break

	case STOP:
		break  
	}

}

func FsmEvTimeOut(){
	switch thisState{
	case DOORSOPEN:
		
		time.Sleep(3000 *time.Milliseconds)
		driver.ElevSetDoorOpenLamp(OFF)

		for i :=0; i > 3; i++ {
			driver.ElevSetButtonLamp(i, currentFloor, OFF)
		}

		direction = queue.QueueChooseDirection(currentFloor, direction)
		driver.ElevSetMotorDirection(direction)
		queue.QueueDeleteCompleted(currentFloor, direction)
		
		
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
		
		driver.ElevSetMotorDirection(DIR_STOP)
		direction = DIR_STOP
		driver.ElevSetStopLamp(ON)
		driver.ElevSetDoorOpenLamp(ON)
		queue.QueueDeleteAllOrders()
	
		thisState = STOP
		break
		
	case MOVING:
		
		driver.ElevSetMotorDirection(DIR_STOP)
		direction = DIR_STOP
		driver.ElevSetStopLamp(ON)
		queue.QueueDeleteAllOrders()
	
		thisState = STOP
		break
		
	case DOORSOPEN:
		
		driver.ElevSetMotorDirection(DIR_STOP)
		direction = DIR_STOP
		driver.ElevSetStopLamp(ON)
		driver.ElevSetDoorOpenLamp(ON)
		queue.QueueDeleteAllOrders()
		
		thisState = STOP
		break
		
	case INIT:
		driver.ElevSetStopLamp(ON)
		if currentFloor == driver.ElevGetFloorSensorSignal(){
			driver.ElevSetDoorOpenLamp(ON)
		}else{
			driver.ElevSetDoorOpenLamp(OFF)
		}

		thisState = STOP
		break
	
	case STOP:
	
		driver.ElevSetStopLamp(ON)
		driver.ElevSetDoorOpenLamp(OFF)
		
		thisState = STOP
		break
	}
}

func FsmEvStopButtonReleased(){
	switch thisState{
	case STOP:

		driver.ElevSetStopLamp(OFF)
		if currentFloor == driver.ElevGetFloorSensorSignal(){
			driver.ElevSetDoorOpenLamp(ON)
		}
		time.Sleep(3000 *time.Milliseconds)
		driver.ElevSetDoorOpenLamp(OFF)
			
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
		if driver.ElevGetFloorSensorSignal() == i {
			driver.ElevSetFloorIndiator(i)
		}
	}
}

func FsmStopButtonisPressed()bool {
	if driver.ElevGetStopSignal() == 1){
		return true
	}else{
		return false
	}
}