package settings

import (
	"flag"
)

var (
	Name       string
	Version    string
	HttpListen string
	GrpcListen string
	fs         *flag.FlagSet
)

func init() {
	Name = NAME
	Version = buildVersion
	fs.StringVar(&HttpListen, "http-listen", "localhost:34000", "bind address and port for http")
	fs.StringVar(&GrpcListen, "grpc-listen", ":3012", "bind address and port for grpc")
}
