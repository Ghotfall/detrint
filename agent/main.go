package main

import (
	"github.com/ghotfall/detrint/agent/server"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	server.Start(server.DefaultHost, server.DefaultPort, logger)
}
