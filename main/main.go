package main

import (
	"fmt"
	"./network"
	"time"
	"sort"
)


func main(){
	ip := network.GetLocalIP()
	fmt.Println("Local ip: ", ip)

	go network.UdpSendAlive()
	peerListCh := make(chan []string)
	go network.UdpRecvAlive(peerListCh)
	sendMsgCh := make(chan network.Msg)
	recvMsgCh := make(chan network.Msg)
	go network.UdpSendMsg(sendMsgCh)
	go network.UdpRecvMsg(recvMsgCh)

	tick := time.Tick(1*time.Second)

	isMaster := false
	peers := []string{ip}

	for {
		select {
		case <-tick:
			sendMsgCh <- network.Msg{5, []bool{true, false}, "129.241.187.150"}

		case peers = <-peerListCh:
			fmt.Println("New peer list: ", peers)
			sort.Strings(peers)
			if len(peers) == 0 {
				fmt.Println("Disconnected, defaulting to master")
				isMaster = true
			} else {
				if peers[0] == ip {
					fmt.Println("We have highest ip, we are master")
					isMaster = true
				} else {
					fmt.Println("We do not have highest ip, we are slave")
					isMaster = false
				}
			}
			fmt.Println("is master: ", isMaster)
			
		case r := <-recvMsgCh:
			fmt.Println("New msg: ", r)
		}
	}


}