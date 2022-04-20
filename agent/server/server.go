package server

import (
	"github.com/ghotfall/detrint/agent/service/shell"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

var DefaultPort = "7056"
var DefaultHost = "0.0.0.0"

func Start(host, port string, l *zap.Logger) {
	address := net.JoinHostPort(host, port)
	listener, listenerErr := net.Listen("tcp", address)
	if listenerErr != nil {
		l.Error("Failed to create listener", zap.String("address", address), zap.Error(listenerErr))
		return
	}

	server := grpc.NewServer()

	// Register services
	shell.Register(server, l)

	l.Info("Starting gRPC server...")
	servErr := server.Serve(listener)
	if servErr != nil {
		l.Error("Failed to start gRPC server", zap.Error(servErr))
		return
	}
}
