package main

import (
	"fmt"

	"micros/rpc/server"
	"micros/settings"
)

func main() {
	settings.Parse()
	fmt.Printf("listen gRPC server %s on %s\n", settings.Version, settings.GrpcListen)
	server.Run(settings.GrpcListen, settings.Debug)
}
