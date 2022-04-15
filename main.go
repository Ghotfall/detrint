package main

import (
	"fmt"
	"github.com/ghotfall/detrint/inv"
	"github.com/ghotfall/detrint/state"
	"go.uber.org/zap"
)

func main() {
	fmt.Println("Project is started!")
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	i := inv.Inventory{
		Machines: map[string]inv.Machine{
			"server_1": {Address: "10.0.0.1"},
			"server_2": {Address: "10.0.0.2"},
			"server_3": {Address: "10.0.0.3"},
			"server_4": {Address: "10.0.0.4"},
		},
		Groups: map[string]inv.Group{
			"group_1": {Members: []string{"server_1", "server_2"}},
			"group_2": {Members: []string{"server_1", "group_1", "server_3"}},
			"group_3": {Members: []string{"group_2", "server_4"}},
		},
	}

	s := state.Set{
		"state 1": state.State{
			Hosts:   "server_1",
			Scripts: []string{"script_1", "script_2"},
		},
		"state 2": state.State{
			Hosts:   "server_2",
			Scripts: []string{"script_2"},
		},
		"state 3": state.State{
			Hosts:   "group_2",
			Scripts: []string{"script_1"},
		},
	}

	s.Deploy(i, logger)
}
