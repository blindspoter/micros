package main

import (
	"fmt"

	"ymir/rpc/server"
	"ymir/settings"
)

func main() {
	settings.Parse()
	fmt.Printf("listen gRPC server %s on %s\n", settings.Version, settings.GrpcListen)
	server.Run(settings.GrpcListen, settings.Debug)
}
