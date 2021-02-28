package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Println("udp连接服务端失败", err)
	}
	defer socket.Close()
	sendData := []byte("hello world!")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("udp消息发送失败", err)
	}
	data := make([]byte, 4096)
	n, remoateAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("udp消息读取失败", err)
		return
	}
	fmt.Println("recv %v, addr %v, count%v \n", string(data[:n]), remoateAddr, n)

}
