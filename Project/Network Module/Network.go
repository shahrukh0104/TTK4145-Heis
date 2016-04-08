package network

import(
	"fmt"
    "strconv"
    "net"
    "log"

)

var laddr, baddr *net.UDPAddr

type UDPmsg struct{
	Raddr string
	Data []byte
	Length int
}


func udp_receive_server(lconn, bconn *net.UDPConn, message_size int, receive_ch chan Udp_message) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in udp_receive_server:", r)
			lconn.Close()
			bconn.Close()
		}
	}()
	bconn_rcv_ch := make(chan Udp_message)
	lconn_rcv_ch := make(chan Udp_message)
	go udp_connection_reader(lconn, message_size, lconn_rcv_ch)
	go udp_connection_reader(bconn, message_size, bconn_rcv_ch)
	for {
		select {
		case buf := <-bconn_rcv_ch:
			receive_ch <- buf
		case buf := <-lconn_rcv_ch:
			receive_ch <- buf
		}
	}
}

func UDPInit(lconn, bconn *net.UDPConn, message_size int, receive_ch chan Udp_message) (err error) {
	udpAddr, err = net.ResolveUDPAddr("udp", "129.241.187.147:"+strconv.Itoa(broadcastListenPort))
	if err != nil {
		log.Fatal(err)
	}
	tempConn, err := net.DialUDP("udp", nil, baddr)
	defer tempConn.Close()
	tempAddr := tempConn.LocalAddr()
	laddr, err = net.ResolveUDPAddr("udp", tempAddr.String())
	laddr.Port = localListenPort
	localListenConn, err := net.ListenUDP("udp4", Laddr)
	if err != nil {
		log.Fatal(err)
	}
	broadcastListenConn, err := net.ListenUDP("udp", baddr)
	if err != nil {
		localListenConn.Close()
		log.Fatal(err)
	}
	go udp_receive_server(localListenConn, broadcastListenConn, message_size, receive_ch)
	go udp_transmit_server(localListenConn, broadcastListenConn, send_ch)
	return err
}

func udp_transmit_server(lconn, bconn *net.UDPConn, send_ch chan Udp_message) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in udp_transmit_server:", r)
			lconn.Close()
			bconn.Close()
		}
	}()
	var err error
	var n int
	for {
		msg := <-send_ch
		if msg.Raddr == "broadcast" {
			n, err = lconn.WriteToUDP(msg.Data, baddr)
		} 
		else {
			raddr, err := net.ResolveUDPAddr("udp", msg.Raddr)
			if err != nil {
				fmt.Printf("Error: udp_transmit_server: could not resolve raddr\n")
				log.Fatal(err)
			}
			n, err = lconn.WriteToUDP(msg.Data, raddr)
		}
		if err != nil || n < 0 {
			fmt.Printf("Error: udp_transmit_server: writing")
			log.Fatal(err)
		}
		
	}
}

func udp_connection_reader(conn *net.UDPConn, message_size int, rcv_ch chan Udp_message) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in udp_connection_reader", r)
			conn.Close()
		}
	}()

	for {
		buf := make([]byte, message_size)
		if err != nil || n < 0 {
			fmt.Printf("Error: udp_connection_reader\n")
			log.Fatal(err)
		}
		rcv_ch <- Udp_message{Raddr: raddr.String(), Data: buf, Length: n}
	}
}
