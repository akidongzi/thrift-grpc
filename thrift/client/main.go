package main

import (
	"context"
	"demo-groc/thrift/th/gen-go/hello_thrift"
	"demo-groc/thrift/th/gen-go/self_thrift"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"net"
	"os"
)

func main() {
	ctx := context.Background()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9888"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	useTransport, err := transportFactory.GetTransport(transport)
	// 这是sayhi
	client := hello_thrift.NewHelloThriftClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:9888", " ", err)
		os.Exit(1)
	}
	defer transport.Close()
	req := hello_thrift.Req{Msg: "我是客户端消息"}
	res, err := client.SayHi(ctx, &req)
	if err != nil {
		log.Println("Echo failed:", err)
		return
	}
	log.Println("response:", res.Msg)


	// self
	selfClient := self_thrift.NewSelfHelloThriftClientFactory(useTransport, protocolFactory)
	sreq := self_thrift.Req{Msg: " 1111111111111"}
	sres ,err := selfClient.SelfSayHi(ctx, &sreq)
	if err != nil {
		log.Println("Echo failed:", err)
		return
	}
	log.Println("response:", sres.Msg)

	fmt.Println("well done")
}
