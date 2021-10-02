package configor

import (
	"fmt"

	"github.com/Summer2021-Project/cli/argv"
	"github.com/Summer2021-Project/web-skeleton/config"
	"github.com/jinzhu/configor"
)

func init() {
	// Conf support YAML, JSON, TOML, Shell Environment
	if err := configor.Load(&config.Config, fmt.Sprintf("%s/../conf/config.yml", argv.Program().Dir)); err != nil {
		panic(err)
	}
}
