package main

import (
	"github.com/Summer2021-Project/cli"
	"github.com/Summer2021-Project/dotenv"
	"github.com/Summer2021-Project/grpc-skeleton/commands"
	_ "github.com/Summer2021-Project/grpc-skeleton/config/configor"
	_ "github.com/Summer2021-Project/grpc-skeleton/config/dotenv"
	_ "github.com/Summer2021-Project/grpc-skeleton/di"
)

func main() {
	cli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(dotenv.Getenv("APP_DEBUG").Bool(false))
	cli.AddCommand(commands.Commands...).Run()
}
