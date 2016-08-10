package main

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"micros/models"
	pb "micros/rpc/proto"
)

type rpcServer struct {
	debug bool
}

func newRpcServer(debug bool) *grpc.Server {
	s := grpc.NewServer()
	glog.Infof("new rpc server with debug %v", debug)
	pb.RegisterMicrosServiceServer(s, &rpcServer{debug})
	return s
}

func (s *rpcServer) GetUserNameById(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	name, err := models.GetUserNameById(in.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{Name: name}, err
}
