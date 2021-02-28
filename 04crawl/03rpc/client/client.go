package main

import (
	"04crawl/03rpc"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, e := net.Dial("tcp", ":1234")
	if e != nil {
		panic(e)
	}
	client := jsonrpc.NewClient(conn)

	var result float64

	e = client.Call("DemoService.Div", _3rpc.Args{10, 0}, &result)

	if e != nil {
		fmt.Printf("%v", e)
	}

	fmt.Println(result)

}
