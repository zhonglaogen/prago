package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var i int

	for {
		i++
		msg, err := decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取数据失败", err)
			break
		}
		fmt.Println(i, "收到的数据为:", msg)
	}

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("服务端端口监听失败", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("连接接受失败", err)
			continue
		}
		go process(conn)
	}

}
