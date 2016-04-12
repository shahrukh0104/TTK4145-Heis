package network

import (
	. "net"
	. "strings"
	. "encoding/json"
	"time"
	"fmt"
)

const udpAlivePort = "12336"
const udpMsgPort = "12334"

func IPString(addr Addr) string {
	return Split(addr.String(), ":")[0]
}

var localIP string

func GetLocalIP() string {

	if localIP == "" {
		addr,_ := ResolveTCPAddr("tcp4","google.com:80")
		localConn, err := DialTCP("tcp4", nil, addr)
		if err == nil{
			localIP = IPString(localConn.LocalAddr())
			localConn.Close()
		}
	}

	return localIP
}

type Msg struct {
	SomeInteger	int
	SomeArray  	[]bool
	Receiver 	string
}

func UdpSendMsg(msgs <-chan Msg){
	localIP := GetLocalIP()
	udpAddr, _ := ResolveUDPAddr("udp4", localIP[0:LastIndex(localIP, ".") + 1]+"255:"+udpMsgPort)
	udpConn, _ := DialUDP("udp4", nil, udpAddr)
	for {
		m := <-msgs
		buf, _ := Marshal(m)
		udpConn.Write(buf)
	}

}

func UdpRecvMsg(msgs chan<- Msg){
	localIP := GetLocalIP()
	udpAddr, _ 		:= ResolveUDPAddr("udp4", ":"+udpMsgPort)
	readConn, _ 	:= ListenUDP("udp4", udpAddr)

	var buf [1024]byte

	for {
		n, fromAddress, _ := readConn.ReadFromUDP(buf[0:])

		if localIP != fromAddress.IP.String() {
			var m Msg
			Unmarshal(buf[0:n], &m)
			if m.Receiver == localIP {
				msgs <- m
			}
		}
	}

}



func UdpSendAlive() {

	localIP := GetLocalIP()
	udpAddr, _ := ResolveUDPAddr("udp4", localIP[0:LastIndex(localIP, ".") + 1]+"255:"+udpAlivePort)
	udpConn, _ := DialUDP("udp4", nil, udpAddr)

	for {
		time.Sleep(15*time.Millisecond)
		udpConn.Write([]byte("I am alive"))
	}
}


func UdpRecvAlive(peerListLocalCh chan []string){

	var buf [1024]byte

	lastSeen := make(map[string]time.Time)
	hasChanges := false
	aliveTimeout := 50*time.Millisecond

	udpAddr, _ 		:= ResolveUDPAddr("udp4", ":"+udpAlivePort)
	readConn, e 	:= ListenUDP("udp4", udpAddr)
	fmt.Println(e)

	for {
		hasChanges = false

		// Ending read after one second has passed
		readConn.SetReadDeadline(time.Now().Add(aliveTimeout))
		_, fromAddress, err := readConn.ReadFromUDP(buf[0:])

		if err != nil {
			continue
		}
		
		addrString := fromAddress.IP.String()

		_, addrIsInList := lastSeen[addrString]

		if !addrIsInList {			
			hasChanges = true
		}

		lastSeen[addrString] = time.Now()

		//Removing IP of dead connection
		for k, v := range lastSeen {
			if time.Now().Sub(v) > aliveTimeout {
				hasChanges = true
				delete(lastSeen, k)
			}
		}
		
		if hasChanges {
			peerList := make([]string, 0, len(lastSeen))

			for k, _ := range lastSeen {
				peerList = append(peerList, k)
			}
			//Sending list of connected IPs
			peerListLocalCh <- peerList
		}
	}
}

