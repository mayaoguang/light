package controller

import (
	"context"
	"fmt"
	"light/example/rpc/example/pb"
	"light/pkg/logging"
)

type HelloController struct{}

// SayHello test router.
func (*HelloController) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	msg := fmt.Sprintf("hello id: %d, name: %s", in.Id, in.Name)
	logging.Infof("SayHello: %s", msg)
	return &pb.HelloReply{Message: msg}, nil
}
