package inv

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

var DefaultPort = "7056"

type Machine struct {
	Address   string                 `toml:"address" json:"address"`
	Username  string                 `toml:"username" json:"username"`
	Password  string                 `toml:"password" json:"password"`
	Variables map[string]interface{} `toml:"variables,omitempty" json:"variables,omitempty"`
}

func (m Machine) Target() (string, error) {
	if len(strings.TrimSpace(m.Address)) == 0 {
		return "", fmt.Errorf("address field is empty, can't make target")
	}

	u := url.URL{Host: m.Address}
	if len(u.Port()) == 0 {
		return net.JoinHostPort(u.Hostname(), DefaultPort), nil
	} else {
		return m.Address, nil
	}
}
