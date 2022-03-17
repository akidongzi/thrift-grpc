package main

import (
	"context"
	hello_grpc "demo-groc/grpc/pb"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	//客户端建立端口信息
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	defer conn.Close()
	if err == nil {
		//客户端建立链接信息
		client := hello_grpc.NewHelloGRPCClient(conn)
		//进行请求发送
		res, _ := client.SayHi(context.Background(), &hello_grpc.Req{Message: "客户端消息"})
		//拿到相应结果
		fmt.Println(res.GetMessage())

	}
}
