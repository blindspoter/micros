package grpcclient

import (
	"log"

	context "golang.org/x/net/context"
	pb "micros/rpc/proto"
)

func GetUserNameById(userID int64) (*pb.UserResponse, error) {
	client := Connect()
	r, err := client.GetUserNameById(context.Background(), &pb.UserRequest{UserId: userID})
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return r, nil
}
