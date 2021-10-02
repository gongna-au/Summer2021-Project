package main

import (
	"github.com/Summer2021-Project/api-skeleton/commands"
	_ "github.com/Summer2021-Project/api-skeleton/config/configor"
	_ "github.com/Summer2021-Project/api-skeleton/config/dotenv"
	_ "github.com/Summer2021-Project/api-skeleton/di"
	"github.com/Summer2021-Project/cli"
	"github.com/Summer2021-Project/dotenv"
)

func main() {
	cli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(dotenv.Getenv("APP_DEBUG").Bool(false))
	cli.AddCommand(commands.Commands...).Run()
}
