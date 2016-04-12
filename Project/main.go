package main

import (
	."../Project/driver"
	."../Project/defines"
	."../Project/fsm"
	"fmt"
)


func main() {

	fmt.Println("Press STOP button to stop elevator and exit program.\n")

	ElevInit(ELEVTYPE_COMEDI)
	floorCh := make(chan int, 1)
	buttonPressCh := make(chan ButtonPress, 1)
	
	go ElevPoller(floorCh, buttonPressCh)
	go Fsm(floorCh, buttonPressCh)

	

	select {}
}







//git add .

//git commit -a -m "Fsm nearly complete"

//git push











