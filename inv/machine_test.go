package inv

import "testing"

func TestMachine_Target(t *testing.T) {
	m := Machine{}
	defaultPort := "3030"

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "all_specified", input: "127.0.0.1:8080", want: "127.0.0.1:8080", wantErr: false},
		{name: "without_port", input: "127.0.0.1", want: "127.0.0.1:" + defaultPort, wantErr: false},
		{name: "empty", input: "   ", want: "", wantErr: true},
		{name: "all_specified_ipv6", input: "[fe80::1%lo0]:53", want: "[fe80::1%lo0]:53", wantErr: false},
		{name: "without_port_ipv6", input: "[fe80::1%lo0]", want: "[fe80::1%lo0]:" + defaultPort, wantErr: false},
		{name: "with_colon_ipv6", input: "[fe80::1%lo0]:", want: "[fe80::1%lo0]:" + defaultPort, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.Address = tt.input

			got, err := m.Target(defaultPort)
			if (err != nil) != tt.wantErr {
				t.Errorf("Target() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Target() got = %v, want %v", got, tt.want)
			}
		})
	}
}
