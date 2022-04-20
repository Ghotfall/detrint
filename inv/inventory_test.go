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

func TestInventory_GetMachineVars(t *testing.T) {
	i := Inventory{
		Machines: map[string]Machine{
			"server_1": {Variables: map[string]interface{}{"var_1": true}},
			"server_2": {Variables: map[string]interface{}{"var_2": 123}},
			"server_3": {Variables: map[string]interface{}{"var_3": "no"}},
			"server_4": {Variables: map[string]interface{}{"var_4": 4.5, "var_41": false}},
		},
		Groups: map[string]Group{
			"group_1": {
				Members: []string{"server_1", "server_2"},
				Variables: map[string]interface{}{
					"gr_var_1": true,
					"gr_var_2": "test",
				},
			},
		},
	}

	tests := []struct {
		name    string
		input   string
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "server_single_var",
			input:   "server_3",
			want:    map[string]interface{}{"var_3": "no"},
			wantErr: false,
		},
		{
			name:    "server_multiple_vars",
			input:   "server_4",
			want:    map[string]interface{}{"var_4": 4.5, "var_41": false},
			wantErr: false,
		},
		{
			name:    "server_group_1",
			input:   "server_1",
			want:    map[string]interface{}{"var_1": true, "gr_var_1": true, "gr_var_2": "test"},
			wantErr: false,
		},
		{
			name:    "server_group_2",
			input:   "server_2",
			want:    map[string]interface{}{"var_2": 123, "gr_var_1": true, "gr_var_2": "test"},
			wantErr: false,
		},
		{
			name:    "no_server",
			input:   "server_5",
			want:    map[string]interface{}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := i.GetMachineVars(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMachineVars() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMachineVars() got = %v, want %v", got, tt.want)
			}
		})
	}
}
