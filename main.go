package main

import (
	"github.com/ghotfall/detrint/inv"
	"github.com/ghotfall/detrint/state"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	i := inv.Inventory{
		Machines: map[string]inv.Machine{
			"server_1": {Address: "127.0.0.1", Variables: map[string]interface{}{"var1": true}},
			"server_2": {Address: "127.0.0.1", Variables: map[string]interface{}{"var2": "test"}},
			"server_3": {Address: "127.0.0.1"},
			"server_4": {Address: "127.0.0.1"},
		},
		Groups: map[string]inv.Group{
			"group_1": {Members: []string{"server_1", "server_2"}, Variables: map[string]interface{}{"gr_var1": 123}},
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
