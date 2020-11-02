package app

import (
	"os"

	"radio/app/config"
)

const (
	EnvProd = "prod"
)

var Dev bool

func Startup() {
	parseDevFlag()
	configureLogger()
	config.Parse()
}

func parseDevFlag() {
	Dev = os.Getenv("RADIO_ENV") != EnvProd
}
