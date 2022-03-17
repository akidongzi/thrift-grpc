package main

import (
	"context"
	hello_grpc "demo-groc/grpc/pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

//服务端hello_grpc.UnimplementedHelloGRPCServer，对应挂载方法的地方
type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

// SayHi 服务端处理请求信息
func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "服务端返回信息"}, nil
}
func main() {
	//建立监听
	lis, _ := net.Listen("tcp", ":8888")
	//服务端生产
	ser := grpc.NewServer()
	//注册信息
	hello_grpc.RegisterHelloGRPCServer(ser, &server{})

	//监听服务端信息
	err := ser.Serve(lis)
	if err != nil {
		return
	}

}
