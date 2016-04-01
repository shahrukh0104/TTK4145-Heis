package driver // where "driver" is the folder that contains io.go, io.c, io.h, channels.go, channels.h and driver.go
/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: -lcomedi -lm
#include "elev.h"
#include "channels.h"
#include "io.h"
*/
import "C"

func Elev_init() int {
	return int(C.elev_init())
}

func Elev_set_speed(speed int) {
	C.elev_set_speed(C.int(speed))
}

func Elev_set_door_open_lamp(value int) {
	C.elev_set_door_open_lamp(C.int(value))
}

func Elev_get_obstruction_signal() int {
	return int(C.elev_get_obstruction_signal())
}

func Elev_get_stop_signal() int {
	return int(C.elev_get_stop_signal())
}

func Elev_set_stop_lamp(value int) {
	C.elev_set_stop_lamp(C.int(value))
}

func Elev_get_floor_sensor_signal() int {
	return int(C.elev_get_floor_sensor_signal())
}

func Elev_set_floor_indicator(floor int) {
	C.elev_set_floor_indicator(C.int(floor))
}

func Elev_get_button_signal(button int, floor int) int {
	return int(C.elev_get_button_signal(C.elev_button_type_t(button), C.int(floor)))
}

func Elev_set_button_lamp(button int, floor int, value int) {
	C.elev_set_button_lamp(C.elev_button_type_t(button), C.int(floor), C.int(value))

}
