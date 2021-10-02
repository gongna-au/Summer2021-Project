package commands

import (
	"github.com/Summer2021-Project/cli"
)

var Commands = []*cli.Command{
	{
		Name:  "grpc:server",
		Short: "gRPC server demo",
		Options: []*cli.Option{
			{
				Names: []string{"d", "daemon"},
				Usage: "Run in the background",
			},
		},
		RunI: &GrpcServerCommand{},
	},
	{
		Name:  "grpc:client",
		Short: "gRPC client demo",
		RunI:  &GrpcClientCommand{},
	},
}
