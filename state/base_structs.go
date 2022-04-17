package state

import (
	"github.com/ghotfall/detrint/inv"
	"github.com/traefik/yaegi/interp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ModulesPath = "./modules"
var ScriptsPath = "./scripts"

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

		for _, machine := range hosts {
			// Get correct target address
			target, targetErr := machine.Target()
			if targetErr != nil {
				l.Error(
					"Can't make target string",
					zap.String("address", machine.Address),
					zap.Error(targetErr),
				)
				continue
			}

			// Create gRPC-connection for client creation
			conn, dialErr := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if dialErr != nil {
				l.Error(
					"Failed to create client connection, skipping target",
					zap.String("target", target),
					zap.Error(dialErr),
				)
				continue
			}

			// Create interpreter
			interpreter := interp.New(interp.Options{
				GoPath:       ModulesPath,
				Unrestricted: true,
			})

			// Execute scripts
			l.Info("Scripts to execute", zap.Strings("scripts", state.Scripts))
			for _, script := range state.Scripts {
				l.Info("Executing script", zap.String("script", script))

				_, evalErr := interpreter.EvalPath(ScriptsPath + "/" + script + ".go")
				if evalErr != nil {
					l.Error(
						"An failure occurred during script execution",
						zap.String("script", script),
						zap.Error(evalErr),
					)
				}
			}

			// Close gRPC-connection
			closeErr := conn.Close()
			if closeErr != nil {
				l.Error(
					"Failed to close client connection",
					zap.String("target", target),
					zap.Error(closeErr),
				)
			}
		}
	}
}
