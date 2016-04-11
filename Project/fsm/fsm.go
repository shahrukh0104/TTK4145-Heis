package fsm

import(
	."./driver"
	."./defines"
	."./timer"
	."./queue"
	"fmt"
)



func Fsm(floorCh chan int, buttonPressCh chan ButtonPress){

	var prevDir = 0
	var doorCloseCh <-chan time.Time

	Msg := MSG{
		INIT,
		-1
		DIR_STOP,
		OrderUp[] 		= {0,0,0,0}
		OrderDown[] 	= {0,0,0,0}
		OrderInside[]	= {0,0,0,0}
		}
	}

	if ElevGetFloorSensorSignal() != -1{
		ElevSetMotorDir(DIR_DOWN)
		prevDir = Dir
		Msg.State = MOVING
	}
	else
		Msg.State = IDLE


	for {
		select {
		case f := <-floorCh:
			if QueueShouldStop(f, PrevDir) == 0{
				ElevSetMotorDir(DIR_STOP)
				ElevSetDoorOpenLamp(LIGHT_ON)
				

				Msg.State = DOORSOPEN
				doorCloseCh = time.After(3*time.Second)
			}
			break

		case <-doorCloseCh:
			
			ElevSetDoorOpenLamp(LIGHT_OFF)
			for i :=0; i > 3; i++ {
                ElevSetButtonLamp(i, floorCh, LIGHT_OFF)
            }
			
			Dir = QueueChooseDirection(floorCh,PrevDir)
			ElevSetMotorDir(Dir)
			QueueDeleteCompleted(floorCh, Dir)
			
			if Dir = DIR_STOP{
				Msg.State = IDLE
			}
			else{
				Msg.State = MOVING
			}
			break

		case b:= <-buttonPressCh:
			switch Msg.State{
			case MOVING:
				QueueAddOrder(floorCh, b)
				break
				
			case IDLE:
				QueueAddOrder(floorCh, b)
				break

			case INIT:
				QueueAddOrder(floorCh, b)
				break

			case DOORSOPEN:
				QueueAddOrder(floorCh, b)
			}
		}
	}
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


