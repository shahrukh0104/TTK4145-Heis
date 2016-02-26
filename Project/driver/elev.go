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














//Hva trenger heisen å vite fra driveren
// Hvilken knapper har blitt trykket inn
// Etasjelys












func elevInit(){
	fmt.Println("Initializing driver")

	lampChannelMatrix = matrixNew(N_FlOORS, N_BUTTONS)
	buttonChannelMatrix = matrixNew(N_FlOORS, N_BUTTONS)
	elevInitMatrices()

	if(!io_init()){
		fmt.Println("Driver initialization failed")
	}

	go clearAllButtonLamps()

	elevDoorOpenLamp(false)
	elevSetStopLamp(false)
	elevSetFloorIndicator(0)


	//Kjøre ned til 1 etasje?
	//Sette opp tråder
	//En struct for buttons (etasje og push)
	fmt.Println("Driver initialization complete")
}


