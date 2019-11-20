package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed", err)
		return
	}
	defer listen.Close()
	for {
		var buf [1024]byte
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("read from client failed", err)
		}
		fmt.Println("recv from", addr, "data:", string(buf[:n]))
		_, err = listen.WriteTo(buf[:n], addr)
		if err != nil {
			fmt.Println("write to client fail", err)
		}
	}

}
