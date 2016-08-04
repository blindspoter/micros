package server

import (
	"log"
	"net"
)

func Run(addr string, debug bool) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}
	s := newRpcServer(debug)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error in Serve: %s", err)
	}
}
