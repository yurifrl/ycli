package types

import (
	log "github.com/charmbracelet/log"
)

type Config struct {
	OpenAIAPIKey string
	log          *log.Logger
}
