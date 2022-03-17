package main

import (
	"context"
	"demo-groc/thrift/th/gen-go/hello_thrift"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

type HelloThriftServer struct{}

func (s *HelloThriftServer) SayHi(ctx context.Context, req *hello_thrift.Req) (res *hello_thrift.Res, err error) {
	fmt.Println("我是来自客户端的消息：", req.GetMsg())
	fmt.Println("我要开始返回消息了：", res.GetMsg())
	return &hello_thrift.Res{Msg: "我是服务端返回的消息"}, nil
}

func main() {
	//定义服务端
	transport, err := thrift.NewTServerSocket(":9888")
	if err != nil {
		panic(err)
	}
	handler := &HelloThriftServer{}
	processor := hello_thrift.NewHelloThriftProcessor(handler)

	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	protocolFactory := thrift.NewTCompactProtocolFactory()
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)
	defer transport.Close()
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
