package main

import (
	"context"
	"demo-groc/thrift/th/gen-go/hello_thrift"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
)

type HelloThriftServer struct {
	hello_thrift.HelloThrift
}

func (s *HelloThriftServer) SayHi(ctx context.Context, req *hello_thrift.Req) (res *hello_thrift.Res, err error) {
	fmt.Println("我是来自客户端的消息：", req.GetMsg())
	fmt.Println("我要开始返回消息了：")
	return &hello_thrift.Res{Msg: "我是服务端返回的消息"}, nil
}

func main() {
	// 定义并启动Thrift服务端
	// 本段代码主要完成了以下步骤：
	// 1. 创建一个监听9888端口的服务器端套接字；
	// 2. 实例化一个处理程序，用于处理HelloThrift的请求；
	// 3. 创建一个处理器，将处理程序与服务端套接字关联；
	// 4. 配置传输工厂和协议工厂，用于序列化和反序列化数据；
	// 5. 使用配置好的处理器、传输工厂、协议工厂启动服务器；
	// 6. 确保服务器启动过程中遇到的任何错误都会导致程序崩溃。
	transport, err := thrift.NewTServerSocket(":9888") // 创建服务器端套接字
	if err != nil {
		panic(err) // 如果创建失败，则直接崩溃
	}
	handler := &HelloThriftServer{} // 实例化处理程序
	processor := hello_thrift.NewHelloThriftProcessor(handler) // 创建处理器

	transportFactory := thrift.NewTBufferedTransportFactory(8192) // 配置传输工厂

	protocolFactory := thrift.NewTCompactProtocolFactory() // 配置协议工厂
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	) // 创建服务器实例
	defer transport.Close() // 确保服务器停止时关闭套接字
	if err := server.Serve(); err != nil { // 启动服务器
		panic(err) // 如果启动失败，则直接崩溃
	}
}
