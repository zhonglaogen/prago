package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("服务端连接异常:", err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') //读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			fmt.Println("消息发送失败", err)
			return
		}
		buf := [521]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("消息接受失败")
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
