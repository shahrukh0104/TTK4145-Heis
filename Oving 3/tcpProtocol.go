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

	fmt.Println("Listening for TCP connection on " + "30000" + ":" + "20016")
	addr, _ := net.ResolveTCPAddr("tcp", PORT)
	tcpListenSocket, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
	defer tcpListenSocket.Close()
	for {
		tcpConnectSocket, err := tcpListenSocket.AcceptTCP()
		defer tcpConnectSocket.Close()

		if err != nil {
			fmt.Println(err)
		}
		go tcpHandleSomethingSomething(tcpConnectSocket)

	}
}

func TCPclient() {
	fmt.Println("Connect to this socket: 129.241.187.147")
	tcpConnectSocket, err := net.Dial("tcp", "129.241.187.147:33546")

	if err != nil {
		fmt.Println(err)
	}

	tcpConnectSocket.Write([]byte("Connect to: 129.241.187.147:33546\x00"))
}

func tcpHandleSomethingSomething(socket *net.TCPConn) {
	fmt.Println("message recieved")
}

func main() {
	TCPserver()
	TCPclient()

}
