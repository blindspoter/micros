package settings

import (
	"flag"
	"github.com/wealthworks/envflagset"
)

var (
	fs         *flag.FlagSet
	Debug      bool
	Name       string
	Version    string
	HttpListen string
	GrpcListen string
)

func init() {
	Name = NAME
	Version = buildVersion
	fs = envflagset.New(NAME, buildVersion)
	fs.BoolVar(&Debug, "debug", false, "app in debug mode")
	fs.StringVar(&HttpListen, "http-listen", "localhost:3011", "bind address and port for http")
	fs.StringVar(&GrpcListen, "grpc-listen", ":3012", "bind address and port for grpc")
}

func Parse() {
	envflagset.Parse()
}
