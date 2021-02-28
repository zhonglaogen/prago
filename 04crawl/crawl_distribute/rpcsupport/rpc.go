package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServerRpc(host string, service interface{}) error {
	rpc.Register(service)

	listener, e := net.Listen("tcp", host)

	if e != nil {
		return e
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

func NewClient(host string) (*rpc.Client, error) {
	conn, e := net.Dial("tcp", host)
	if e != nil {
		return nil, e
	}
	return jsonrpc.NewClient(conn), nil

}
