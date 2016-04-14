package network

import (
	. "encoding/json"
	"fmt"
	"sort"
	"time"
)

var NEWORDER int = 0
var COMPLETEDORDER int = 1
var ELEVATORSTATE int = 2

type NewOrder struct {
	Floor      int
	ButtonType int
	IP         string
}

type CompletedOrder struct {
	Floor      int
	ButtonType int
}

type ElevatorState struct {
	Floor        int
	Dir          int
	OrdersInside []int
	State        int
}

func NetworkMsgPacker(msgToNetworkCh chan<- NetWorkMsg, newOrderCh <-chan NewOrder, CompletedOrderCh <-chan CompletedOrder, elevatorStateCh <-chan ElevatorState) {
	for {
		select {
		case newOrder := <-newOrderCh:
			buf := json.Marshal(newOrder)
			msgToNetworkCh <- NetWorkMsg{Msgtype: NEWORDER, Data: string(buf)}

		case completedOrder := <-CompletedOrderCh:
			buf := json.Marshal(completedOrder)
			msgToNetworkCh <- NetWorkMsg{Msgtype: COMPLETEDORDER, Data: string(buf)}

		case elevatorState := <-elevatorStateCh:
			buf := json.Marshal(elevatorState)
			msgToNetworkCh <- NetWorkMsg{Msgtype: ELEVATORSTATE, Data: string(buf)}
		}
	}

}
func NetworkMsgUnpacker(msgFromNetworkCh <-chan NetWorkMsg, newOrderCh chan<- NewOrder, CompletedOrderCh chan<- CompletedOrder, ElevatorStateCh chan<- ElevatorState) {
	for {
		msg := <-msgFromNetworkCh
		switch msg.Msgtype {
		case NEWORDER:
			var newOrder NewOrder
			json.Unmarshal(msg.Data, &newOrder)
			newOrderCh <- newOrder

		case COMPLETEDORDER:
			var completedOrder CompletedOrder
			json.Unmarshal(msg.Data, &completedOrder)
			completedOrderCh <- completedOrder

		case ELEVATORSTATE:
			var elevatorState ElevatorState
			json.Unmarshal(msg.Data, &elevatorState)
			elevatorStateCh <- elevatorState

		}
	}
}

/*
func Network() {
	ip := GetLocalIP()
	fmt.Println("Local ip: ", ip)
	/*
		go UdpSendAlive()
		peerListCh := make(chan []string)
		go UdpRecvAlive(peerListCh)
		sendMsgCh := make(chan MSG)
		recvMsgCh := make(chan MSG)
		go UdpSendMsg(sendMsgCh)
		go UdpRecvMsg(recvMsgCh)

	tick := time.Tick(1 * time.Second)

	isMaster := false
	peers := []string{ip}

	for {
		select {
		case <-tick:
			//sendMsgCh <- network.Msg{5, []bool{true, false}, "129.241.187.157"}
			sendMsgCh <- MSG{}

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
*/
