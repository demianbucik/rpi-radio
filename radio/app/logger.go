package app

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
)

func configureLogger() {
	if Dev {
		log.SetHandler(text.Default)
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetHandler(json.Default)
		log.SetLevel(log.InfoLevel)
	}
}
