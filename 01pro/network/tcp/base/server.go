package main

import (
	"bufio"
	"fmt"
	"net"
)

func processs(conn net.Conn) {
	defer conn.Close() //关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("read from client fail err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到的数据:", recvStr)
		sendStr := "来自服务端的数据:" + recvStr
		conn.Write([]byte(sendStr)) //发送数据
	}
}

func main() {
	var listern, err = net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen fail err", err)
		return
	}
	for {
		// 建立连接
		conn, err := listern.Accept()
		if err != nil {
			fmt.Println("accept fail, err:", err)
			continue
		}
		go processs(conn)
	}
}
