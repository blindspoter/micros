package main

import (
	"fmt"

	"micros/settings"
)

func main() {
	settings.Parse()
	fmt.Printf("listen gRPC server %s on %s\n", settings.Version, settings.GrpcListen)
	Run(settings.GrpcListen, settings.Debug)
}
