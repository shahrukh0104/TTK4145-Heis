package driver


import(
	"fmt"
	"net"
	"Matrix"
	"io"
)

const (
	N_FLOORS = 4
	N_BUTTONS = 3
)

var lampChannelMatrix Matrix
var buttonChannelMatrix Matrix










func elev_set_stop_lamp() {
	
}





func elev_get_stop_signal() int {
	return io_read_bit(STOP)
}







func elev_get_obstruction_signal() int {
	return io_read_bit(OBSTRUCTION)
}


















func elev_set_door_open_lamp(value int){
	if (value != 0){
        io_set_bit(LIGHT_DOOR_OPEN)
	}
	else{
        io_clear_bit(LIGHT_DOOR_OPEN)
    }
}
















func elev_set_speed(speed int, lastSpeed int){

    // If to start (speed > 0)
    if (speed > 0){
        io_clear_bit(MOTORDIR)
    }

    else if (speed < 0){
        io_set_bit(MOTORDIR)
    }

    // If to stop (speed == 0)
    else if (last_speed < 0){
        io_clear_bit(MOTORDIR)
    }

    else if (last_speed > 0){
        io_set_bit(MOTORDIR);
    }

    last_speed := speed
    
    // Write new setting to motor.
    io_write_analog(MOTOR, 2048 + 4 * abs(speed));
}


















elev_get_floor_sensor_signal
elev_set_floor_indicator
elev_get_button_signal
elev_set_button_lamp












func elevClearAllButtonLamps(){
for i := 0; i < N_FLOORS; i++ {
        if (i != 0)
            elevSetButtonLamp(BUTTON_CALL_DOWN, i, 0);

        if (i != N_FLOORS - 1)
            elevSetButtonLamp(BUTTON_CALL_UP, i, 0);

        elevSetButtonLamp(BUTTON_COMMAND, i, 0);
    }

}












func elevInitMatrices() {
	lampChannelMatrix.Set(0,0,LIGHT_UP1)
	lampChannelMatrix.Set(0,1,LIGHT_DOWN1)
	lampChannelMatrix.Set(0,2,LIGHT_COMMAND1)
	
	lampChannelMatrix.Set(1,0,LIGHT_UP2)
	lampChannelMatrix.Set(1,1,LIGHT_DOWN2)
	lampChannelMatrix.Set(1,2, LIGHT_COMMAND2)
	
	lampChannelMatrix.Set(2,0,LIGHT_UP3)
	lampChannelMatrix.Set(2,1,LIGHT_DOWN3)
	lampChannelMatrix.Set(2,2,LIGHT_COMMAND3)
	
	lampChannelMatrix.Set(3,0,LIGHT_UP4)
	lampChannelMatrix.Set(3,1,LIGHT_DOWN4)
	lampChannelMatrix.Set(3,2,LIGHT_COMMAND4)
	
	
	buttonChannelMatrix.Set(0,0,FLOOR_UP1)
	buttonChannelMatrix.Set(0,1,FLOOR_DOWN1)
	buttonChannelMatrix.Set(0,2,FLOOR_COMMAND1)
	
	buttonChannelMatrix.Set(1,0,FLOOR_UP2)
	buttonChannelMatrix.Set(1,1,FLOOR_DOWN2)
	buttonChannelMatrix.Set(1,2,FLOOR_COMMAND2)
	
	buttonChannelMatrix.Set(2,0,FLOOR_UP3)
	buttonChannelMatrix.Set(2,1,FLOOR_DOWN3)
	buttonChannelMatrix.Set(2,2,FLOOR_COMMAND3)
	
	buttonChannelMatrix.Set(3,0,FLOOR_UP4)
	buttonChannelMatrix.Set(3,1,FLOOR_DOWN4)
	buttonChannelMatrix.Set(3,2,FLOOR_COMMAND4)
}


















func elevInit(){
	fmt.Println("Initializing driver")

	lampChannelMatrix = matrixNewMatrix(N_FlOORS, N_BUTTONS)
	buttonChannelMatrix = matrixNewMatrix(N_FlOORS, N_BUTTONS)
	elevInitMatrices()

	if(!io_init()){
		fmt.Println("Driver initialization failed")
	}

	go clearAllButtonLamps()

	elevDoorOpenLamp(false)
	elevSetStopLamp(false)
	elevSetFloorIndicator(0)


	//Kjøre ned til 1 etasje?
	fmt.Println("Driver initialization complete")
}


