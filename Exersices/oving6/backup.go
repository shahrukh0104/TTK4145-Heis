package main

import(
	"log"
	"net"
	"time"
	"encoding/binary"
	"os/exec"
)

func primary(start int, udpBroadcast *net.UDPConn){

	newBackup := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run backup.go")
	err := newBackup.Run()
	if err != nil {
		log.Fatal(err)
	}

	msg := make([]byte, 1)

	for i := start;; i++{
		log.Println(i)
		msg[0] = byte(i);
		udpBroadcast.Write(msg)
		time.Sleep(100*time.Millisecond)
	}
	
}

func backup(udpListen *net.UDPConn) int{
	listenChan := make(chan int, 1); 
	backupvalue := 0
	go listen(listenChan, udpListen)
	for {
		select {
			case backupvalue = <- listenChan:
				time.Sleep(50*time.Millisecond)
				break
			case <-time.After(1*time.Second):
				log.Println("The primary is dead, long live the primary")
				return backupvalue
		}
	}
	
	
}

func listen(listenChan chan int, udpListen *net.UDPConn) {

	buffer := make([]byte, 1024)

	for {
		udpListen.ReadFromUDP(buffer[:])
		if err != nil {
			log.Fatal(err)
		} 
		listenChan <- int(binary.LittleEndian.Uint64(buffer)) //konverterer byteearray til int tror jeg
		time.Sleep(100*time.Millisecond)
	}
	
}

func main() {
	
	udpAddr, err := net.ResolveUDPAddr("udp", ":20016")
	if err != nil {
		log.Fatal(err)
	}
	udpListen, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	
	backupvalue := backup(udpListen)	
	udpListen.Close()	
	udpAddr, err = net.ResolveUDPAddr("udp","129.241.187.147:20016")
	if err != nil {l
		og.Fatal(err)
	}
	udpBroadcast, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {log.Fatal(err)}
	primary(backupvalue, udpBroadcast)
	udpBroadcast.Close()
	
	
}