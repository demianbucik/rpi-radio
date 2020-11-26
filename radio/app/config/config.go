package config

import (
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/apex/log"
)

const (
	configFile = "env.toml"
)

var (
	Env Environ
)

func Parse() {
	configDir := os.Getenv("CONF_DIR")
	if configDir == "" {
		configDir = "./config"
	}
	if _, err := toml.DecodeFile(path.Join(configDir, configFile), &Env); err != nil {
		log.WithError(err).Fatal("Failed to parse config")
	}
}
