package viper

import (
	"fmt"

	"github.com/Summer2021-Project/cli/argv"
	"github.com/Summer2021-Project/web-skeleton/config"
	"github.com/spf13/viper"
)

func init() {
	// Conf support JSON, TOML, YAML, HCL, INI, envfile
	viper.AddConfigPath(fmt.Sprintf("%s/../conf/config.yml", argv.Program().Dir))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config.Config); err != nil {
		panic(err)
	}
}
