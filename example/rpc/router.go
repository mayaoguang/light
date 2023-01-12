package main

import (
	"google.golang.org/grpc"
	"light/example/rpc/example/controller"
	"light/example/rpc/example/pb"
)

// Index router.
func Index(s *grpc.Server) {
	hc := &controller.HelloController{}
	pb.RegisterHelloWorldServiceServer(s, hc)
}
