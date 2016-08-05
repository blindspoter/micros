package main

import (
	"flag"
	"log"
	"net"
)

func Run(addr string, debug bool) {
	flag.Parse()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panicf("failed to listen: %v", err)
	}
	s := newRpcServer(debug)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error in Serve: %s", err)
	}
}
