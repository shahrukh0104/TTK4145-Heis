package driver // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.h and driver.go
/*
#cgo CFLAGS: -std=gnu11
#cgo LDFLAGS: -lcomedi -lm
#include "elev.h"
#include "channels.h"
#include "io.h"
*/
import "C"

func ElevInit(elevatorType int) {
	C.elev_init(C.elev_type(elevatorType))
}

func ElevSetMotorDir(dir int) {
	C.elev_set_motor_direction(C.elev_motor_direction_t(dir))
}

func ElevSetDoorOpenLamp(value int) {
	C.elev_set_door_open_lamp(C.int(value))
}

func ElevGetObstructionSignal() int {
	return int(C.elev_get_obstruction_signal())
}

func ElevGetStopSignal() int {
	return int(C.elev_get_stop_signal())
}

func ElevSetStopLamp(value int) {
	C.elev_set_stop_lamp(C.int(value))
}

func ElevGetFloorSensorSignal() int {
	return int(C.elev_get_floor_sensor_signal())
}

func ElevSetFloorIndicator(floor int) {
	C.elev_set_floor_indicator(C.int(floor))
}

func ElevGetButtonSignal(button int, floor int) int {
	return int(C.elev_get_button_signal(C.elev_button_type_t(button), C.int(floor)))
}

func ElevSetButtonLamp(button int, floor int, value int) {
	C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value))

}

type ButtonPress struct {
	Floor  int
	Button int
}

func ElevPoller(floorChan chan int, buttonPressChan chan ButtonPress) {

	prevFloor := -1
	var prevButtons [4][3]int

	for {

		floor := ElevGetFloorSensorSignal()
		if floor != prevFloor && floor != -1 {
			floorChan <- floor
		}
		prevFloor = floor

		for f := 0; f < 4; f++ {
			for b := 0; b < 3; b++ {
				v := ElevGetButtonSignal(b, f)
				if v != prevButtons[f][b] && v != 0 {
					buttonPressChan <- ButtonPress{f, b}
				}
				prevButtons[f][b] = v
			}
		}

	}
}
