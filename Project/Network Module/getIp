package network

import(
	"net"
	"fmt"
)

type Ip string

func getmyip()Ip{
	interfaces, err := net.Interfaces()
	if err!=nil{
		fmt.Println(err)
	}
	for _,i:= range interfaces{
		addrs, err := i.Addrs()
		if err!=nil{
			fmt.Println(err)
		}
		for _,addr:=range addrs{
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return Ip(ipnet.IP.String()) 
		}
	}
	return 

}

func getsendersip(sender *net.UDPAddr)Ip{
	return Ip(sender.IP.String())
}

