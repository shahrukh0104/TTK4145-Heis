package timer

import(
	"fmt"
	"time"
)

const(
	DoorOpenTime = 3*time.Second
)

var(
	timer time.Time
	timerFlag bool = false
)

func TimeCheck(timeOut chan int) {
	for {
		if (time.Since(timer) > DoorOpenTime) && timerFlag == true {
			timerFlag = false
			timeOut <- 0
			fmt.Println("Timeout")
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func TimerDoor(timer chan string, timeOut chan int) {
	go checkTimer(timeOut)

	for {
		select {
			case <- timer:
				fmt.Println("Starting timer")
				timer = time.Now()
				timerFlag = true
			case <- timeOut:
				timer <- "timeout"
		}
	}
}
