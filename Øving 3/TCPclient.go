package main

import(
        "fmt"
        "net"
)

const(
        server = 
        port_fixed =
        port_
)

func connectTCP(){
    conn, err :=  
    if err != nil{
    	fmt.Println("Couldn't connect to TCP server")
    }
    addr, err := net.ResolveTCPAddr("tcp",server)
    if err != nil{
    	fmt.Println("Couldn't resolve addres for" + port_fixed)    	
    }
    listener, err := net.ListenTCP("tcp",addr)


}
