package api

import (
	"go-seed/pkg/config"
	"go-seed/pkg/logger/zap"
)

func Start(config *config.Configuration) error {
	log := zap.New(config.Logging)
	for {
		log.Info("initialized logger successfully")
	}

	return nil
}