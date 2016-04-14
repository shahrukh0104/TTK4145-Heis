package none

import(
	"fmt"
    "strconv"
    "net"
    "log"

)

var laddr, baddr *net.UDPAddr

type Udp_message struct{
	Raddr string
	Data []byte
	Length int
}

func UdpInit(llport, blport, message_size int, send_ch, receive_ch chan Udp_message) (err error) {
	baddr, err = net.ResolveUDPAddr("udp", "129.241.187.147:"+strconv.Itoa(blport))
	if err != nil {
		return err
	}
	tempConn, err := net.DialUDP("udp", nil, baddr)
	defer tempConn.Close()
	tempAddr := tempConn.LocalAddr()
	laddr, err = net.ResolveUDPAddr("udp", tempAddr.String())
	//laddr.Port = lPort
	llconn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return err
	}
	blconn, err := net.ListenUDP("udp", baddr)
	if err != nil {
		llconn.Close()
		return err
	}
	go UdpReceiveServer(llconn, blconn, message_size, receive_ch)
	go UdpTransmitServer(llconn, blconn, send_ch)
	return err
}


func UdpReceiveServer(lconn, bconn *net.UDPConn, message_size int, receive_ch chan Udp_message) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in receive:", r)
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



func UdpTransmitServer(lconn, bconn *net.UDPConn, send_ch chan Udp_message) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in transmit:", r)
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
		} else {
			raddr, err := net.ResolveUDPAddr("udp", msg.Raddr)
			if err != nil {
				fmt.Printf("Error in transmit\n")
				log.Fatal(err)
			}
			n, err = lconn.WriteToUDP(msg.Data, raddr)
		}
		if err != nil || n < 0 {
			fmt.Printf("Error in transmit")
			log.Fatal(err)
		}
		
	}
}

func udp_connection_reader(conn *net.UDPConn, message_size int, rcv_ch chan Udp_message)(err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in connection", r)
			conn.Close()
		}
	}()
	for {
		buf := make([]byte, message_size)
		n, raddr, err := conn.ReadFromUDP(buf)
		if err != nil || n < 0 {
			fmt.Printf("Error in connection\n")
			log.Fatal(err)
		}
		rcv_ch <- Udp_message{Raddr: raddr.String(), Data: buf, Length: n}
	}
}
