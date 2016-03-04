package main

import (
	"log"
	"net"
	"time"
)

func primary(start int) {
	udpAddr, err := net.ResolveUDPAddr("udp", "129.241.187.147:20016")
	if err != nil {
		log.Fatal(err)
	}

	udpBroadcast, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer udpBroadcast.Close()

	msg := make([]byte, 8)

	for i := start; ; i++ {
		log.Println(i)
		msg[0] = byte(i)
		udpBroadcast.Write(msg)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	primary(1)
}
