package config

import (
	"github.com/BurntSushi/toml"
	"github.com/apex/log"
)

const (
	configFile = "config/env.toml"
)

var (
	Env Environ
)

func Parse() {
	if _, err := toml.DecodeFile(configFile, &Env); err != nil {
		log.WithError(err).Fatal("Failed to parse config")
	}
}
