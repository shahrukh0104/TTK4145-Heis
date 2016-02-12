package main

import (
	"fmt"
	"net"
)

const (
	HOST = "20016"
	PORT = "33546"
)

func TCPserver() {

	fmt.Println("Listening for TCP connection on " + HOST)
	addr, _ := net.ResolveTCPAddr("tcp", ":"+HOST)
	tcpListenSocket, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		defer tcpListenSocket.Close()
		for {
			tcpConnectSocket, err := tcpListenSocket.AcceptTCP()
			defer tcpConnectSocket.Close()
			fmt.Println("Accepted!")
			if err != nil {
				fmt.Println(err)
			}
			go tcpHandleSomethingSomething(tcpConnectSocket)
		}
	}()
}

func TCPclient() {
	fmt.Println("Connecting to this socket: 129.241.187.23")
	tcpConnectSocket, err := net.Dial("tcp", "129.241.187.23:33546")

	if err != nil {
		fmt.Println(err)
		return
	}

	tcpConnectSocket.Write([]byte("Connect to: 129.241.187.147:" + HOST + "\x00"))
	for {
	}
}

func tcpHandleSomethingSomething(conn *net.TCPConn) {
	message := make([]byte, 1024)
	conn.Read(message)
	fmt.Println("Message: " + string(message))

}

func main() {
	TCPserver()
	TCPclient()

}
