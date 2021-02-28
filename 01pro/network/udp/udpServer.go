package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("udp端口监听失败", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("数据读取失败", err)
			continue
		}
		fmt.Printf("data %v, addr %v, count %v", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:], addr)
		if err != nil {
			fmt.Println("数据发送失败", err)
			continue
		}
	}
}
