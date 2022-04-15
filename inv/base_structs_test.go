package inv

import (
	"reflect"
	"testing"
)

func TestInventory_ResolveHosts(t *testing.T) {
	i := Inventory{
		Machines: map[string]Machine{
			"server_1": {Address: "10.0.0.1"},
			"server_2": {Address: "10.0.0.2"},
			"server_3": {Address: "10.0.0.3"},
			"server_4": {Address: "10.0.0.4"},
		},
		Groups: map[string]Group{
			"group_1": {Members: []string{"server_1", "server_2"}},
			"group_2": {Members: []string{"server_1", "group_1", "server_3"}},
			"group_3": {Members: []string{"group_2", "server_4"}},
		},
	}

	tests := []struct {
		name     string
		selector string
		want     map[string]Machine
	}{
		{name: "simple_server", selector: "server_2", want: map[string]Machine{
			"server_2": i.Machines["server_2"],
		}},
		{name: "simple_group", selector: "group_1", want: map[string]Machine{
			"server_1": i.Machines["server_1"],
			"server_2": i.Machines["server_2"],
		}},
		{name: "complex_group", selector: "group_2", want: map[string]Machine{
			"server_1": i.Machines["server_1"],
			"server_2": i.Machines["server_2"],
			"server_3": i.Machines["server_3"],
		}},
		{name: "complex_group_multiple", selector: "group_3", want: map[string]Machine{
			"server_1": i.Machines["server_1"],
			"server_2": i.Machines["server_2"],
			"server_3": i.Machines["server_3"],
			"server_4": i.Machines["server_4"],
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := i.ResolveHosts(tt.selector); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveHosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
