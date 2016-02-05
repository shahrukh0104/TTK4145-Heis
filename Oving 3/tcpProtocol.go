package main

import (
	"fmt"
	"net"
	"time"
)


const (
	HOST = 20016
	PORT = 34933
)


func TCPserver() {

	fmt.Println("Listening for TCP connection on " + HOST + ":" + PORT)
	addr, _ := net.ResolveUDPAddr("udp", ":"+ PORT)
	tcpListenSocket, err := net.ListenTCP("tcp", addr)
	if err != nil{
		fmt.Println(err)
	}
	defer tcpListenSocket.Close()
	for {
		tcpConnectSocket , err := tcpListenSocket.AcceptTCP()
		defer tcpConnectSocket.Close()

		if err != nil{
			fmt.Println(err)
		}
		go tcpHandleSomethingSomething(tcpConnectSocket)	

	}
}






func TCPclient() {
	fmt.Println("Connect to this socket: 129.241.187.147")
	addr, err := net.ResolveTCPAddr("tcp", 129.241.187.147:20016)
	tcpConnectSocket, err := net.DialTCP("tcp", addr)


	
}