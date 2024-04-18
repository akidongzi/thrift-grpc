package main

import (
	"context"
	"demo-groc/thrift/th/gen-go/hello_thrift"
	"demo-groc/thrift/th/gen-go/self_thrift"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
)

type HelloThriftHandler struct{}

func (s *HelloThriftHandler) SayHi(ctx context.Context, req *hello_thrift.Req) (res *hello_thrift.Res, err error) {
	fmt.Println("我是来自 HelloThrift 的消息：", req.GetMsg())
	fmt.Println("我要开始返回消息了：")
	return &hello_thrift.Res{Msg: "我是 HelloThrift 服务端返回的消息"}, nil
}

type SelfHelloThriftHandler struct{}

func (s *SelfHelloThriftHandler) SelfSayHi(ctx context.Context, req *self_thrift.Req) (res *self_thrift.Res, err error) {
	fmt.Println("我是来自 AnotherThrift 的消息：", req.GetMsg())
	fmt.Println("我要开始返回消息了：")
	return &self_thrift.Res{Msg: "我是 AnotherThrift 服务端返回的消息"}, nil
}

func main() {
	transport, err := thrift.NewTServerSocket(":9888")
	if err != nil {
		panic(err)
	}

	// 创建多路复用处理器
	multiplexedProcessor := thrift.NewTMultiplexedProcessor()

	// 实例化 HelloThrift 的处理程序并注册到多路复用处理器中
	helloHandler := &HelloThriftHandler{}
	helloProcessor := hello_thrift.NewHelloThriftProcessor(helloHandler)
	multiplexedProcessor.RegisterProcessor("HelloThrift", helloProcessor)

	// 实例化 AnotherThrift 的处理程序并注册到多路复用处理器中
	anotherHandler := &SelfHelloThriftHandler{}
	anotherProcessor := self_thrift.NewSelfHelloThriftProcessor(anotherHandler)
	multiplexedProcessor.RegisterProcessor("SelfHelloThrift", anotherProcessor)

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	server := thrift.NewTSimpleServer4(
		multiplexedProcessor,
		transport,
		transportFactory,
		protocolFactory,
	)
	defer transport.Close()
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
