package main

import (
	"04crawl/03rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main()  {
	rpc.Register(_3rpc.DemoService{})

	listener, e := net.Listen("tcp", ":1234")

	if e != nil {
		panic(e)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error is %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)



	}

}
