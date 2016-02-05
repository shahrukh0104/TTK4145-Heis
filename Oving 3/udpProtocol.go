package main

import (
	"fmt"
	"net"
	"time"
)

const (
	host    = 20016
	udpPort = "30000" //Listen to this port
)

func udpRecieve(port string) {
	buff := make([]byte, 1024)
	addr, _ := net.ResolveUDPAddr("udp", ":"+port)
	sock, _ := net.ListenUDP("udp", addr) //Listner
	for {
		n, recvaddr, err := sock.ReadFromUDP(buff)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buff[:n]))
		fmt.Printf("%+v\n", recvaddr)
		udpSend(recvaddr)
		time.Sleep(1000 * time.Millisecond)

	}
	defer sock.Close()
}

func udpSend(recvaddr *net.UDPAddr) {
	recvaddr.Port = host

	conn, err := net.DialUDP("udp", nil, recvaddr)
	defer conn.Close()
	if err != nil {
		fmt.Println("Error")
	}

	msg := []byte("Shahrukh Khan\x00")
	conn.Write(msg)
	fmt.Println("Message sent", string(msg))

}

func main() {
	udpRecieve(udpPort)

}
