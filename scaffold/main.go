package main

import (
	"github.com/Summer2021-Project/cli"
	"github.com/Summer2021-Project/dotenv"
	"github.com/Summer2021-Project/scaffold/commands"
)

func main() {
	cli.SetName("Dubbo-gocli").
		SetVersion(commands.CLIVersion).
		SetDebug(dotenv.Getenv("APP_DEBUG").Bool(false))
	cli.AddCommand(commands.Cmds...).Run()
}
