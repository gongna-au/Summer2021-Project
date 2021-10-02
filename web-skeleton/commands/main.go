package commands

import (
	"github.com/Summer2021-Project/cli"
)

var Commands = []*cli.Command{
	{
		Name:  "web",
		Short: "\tStart the web server",
		Options: []*cli.Option{
			{
				Names: []string{"a", "addr"},
				Usage: "\tListen to the specified address",
			},
			{
				Names: []string{"d", "daemon"},
				Usage: "\tRun in the background",
			},
		},
		RunI: &WebCommand{},
	},
}
