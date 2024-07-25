package main

import (
	"ecommerce/configs"
	"ecommerce/pkg/logger"
)

func main() {
	cfg := configs.Load()

	log := logger.NewLogger(cfg.ServiceName, logger.LevelDebug)
	defer logger.Cleanup(log)

}
