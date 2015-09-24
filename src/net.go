package main

import (
	"fmt"
	"net"
)

//获取随机端口
func getPort() int {
	laddr := net.TCPAddr{IP:net.IPv4(127,0,0,1),Port:0}
	l,err := net.ListenTCP("tcp4", &laddr)
	if err == nil {
		addr := l.Addr()
		return addr.(*net.TCPAddr).Port
	}else {
		return 0
	}
}

func main() {
	fmt.Println(getPort())
}