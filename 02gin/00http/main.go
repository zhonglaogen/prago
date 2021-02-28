package main

import (
	"02gin/00http/example/myio/mysock/mynet"
	"02gin/00http/example/myio/mysock/mynet/myhttp2"
	"02gin/00http/example/myio/mysock/mynet/myssl"
	myupload "02gin/00http/example/myio/mysock/mynet/myupload"
	"02gin/00http/example/myio/mysock/mynet/myws"
	"02gin/00http/example/myio/mysock/mynet/template"
)

func main() {

	//=========命令行参数================
	//mycmdline.Mycmdline()

	//=========网络io================
	//mytcp.MyTcpServer()
	//mytcp.MyTcpClient()
	//onetoone.MyUdpServer()
	//onetoone.MyUdpCli()
	//multi.MyMultiServer()
	//multi.MyMultiCli()
	//broadcast.MyBCServer()
	//broadcast.MyBCCli()

	//=========****网络io-http相关================
	mynet.MyHttp1()
	mynet.MyHttp2()
	mynet.MyHttp21()
	mynet.MyHttp22()
	mynet.MyHttp3()
	mynet.MyHttp31()
	mynet.Myhttpcli()
	//
	template.TestTemplate()
	//
	myupload.MyUpServer()
	myupload.MyUpCli()
	////
	myssl.MySsl()
	myssl.MySslCli()
	////
	myhttp2.MyHttp2Server()
	//myhttp2.MyHttp2Cli()
	//
	myws.MyWsServer()
}
