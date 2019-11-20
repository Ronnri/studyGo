package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	c, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		return
	}
	defer c.Close()

	input := bufio.NewReader(os.Stdin)
	for {
		s, err := input.ReadString('\n')
		if err != nil {
			return
		}
		_, err = c.Write([]byte(s))
		if err != nil {
			return
		}

		var buf [1024]byte
		n, addr, err := c.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("read from server failed", err)
			return
		}
		fmt.Println("read from", addr, "data:", string(buf[:n]))

	}

}
