package state

import (
	"github.com/ghotfall/detrint/inv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
)

type State struct {
	Hosts   string   `toml:"hosts"`
	Scripts []string `toml:"scripts"`
}

type Set map[string]State

func (s Set) Deploy(i inv.Inventory, l *zap.Logger) {
	for stateName, state := range s {
		l.Info("Starting state deployment", zap.String("state", stateName))

		hosts := i.ResolveHosts(state.Hosts)
		l.Info("Resolved host", zap.Any("hosts", hosts))

		for host, machine := range hosts {
			addr := machine.Address
			if !strings.Contains(addr, ":") {
				addr = addr + ":7056" // TODO: separate default value
			}

			conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				l.Error("Failed to create client connection", zap.String("target", addr), zap.Error(err))
			}
		}
	}
}
