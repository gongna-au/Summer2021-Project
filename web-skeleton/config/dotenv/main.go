package dotenv

import (
	"fmt"

	"github.com/Summer2021-Project/cli/argv"
	"github.com/Summer2021-Project/dotenv"
)

func init() {
	// Env
	if err := dotenv.Load(fmt.Sprintf("%s/../.env", argv.Program().Dir)); err != nil {
		panic(err)
	}
}
