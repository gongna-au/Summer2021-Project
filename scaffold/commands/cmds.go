package commands

import "github.com/Summer2021-Project/cli"

var Cmds = []*cli.Command{
	{
		Name:  "new",
		Short: "\tCreate a project",
		RunI:  &NewCommand{},
	},
}
