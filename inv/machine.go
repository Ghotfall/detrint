package inv

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

type Machine struct {
	Address   string                 `toml:"address"`
	Username  string                 `toml:"username"`
	Password  string                 `toml:"password"`
	Variables map[string]interface{} `toml:"variables,omitempty"`
}

func (m Machine) Target(defaultPort string) (string, error) {
	if len(strings.TrimSpace(m.Address)) == 0 {
		return "", fmt.Errorf("address field is empty, can't make target")
	}

	u := url.URL{Host: m.Address}
	if len(u.Port()) == 0 {
		return net.JoinHostPort(u.Hostname(), defaultPort), nil
	} else {
		return m.Address, nil
	}
}
