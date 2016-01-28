package main 

import(
        "fmt"
        "net"
        "time
)

const(
        host =
        udpPort = 
        portT =
)

func udpRecieve(port string){
    buff := make([]byte, 1024)
    addr := net.ResolveUDPAddr("udp",":" + port)
    sock := net.ListenUDP("udp",addr)
    for {
        //sende shiten da
        _,_, err:=sock.ReadFromUDP(buff)
		if err != nil {
			fmt.Println(err)
		} 
		fmt.Println(string(buff[:]))
    }
    
}

func udpSend(){
    raddr := net.ResolveUDPAddr("udp",net.JoinHostPort(host,udpPort))
    if err != nil {
        fmt.Println("Failed to get adress for port: " + udpPort)
    }
    conn, err := net..DialUDP("udp", nil, raddr)
    if err != nil{
        fmt.Println("Error")
    }
    go udpReceive(udpPort)
	for{
		time.Sleep(1000*time.Millisecond)
		conn.Write([]byte("Shahrukh Khan"))
		fmt.Println("Message sent")
        }	
}

func main(){
    udpSend()
    
}
