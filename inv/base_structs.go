package inv

type Machine struct {
	Address   string                 `toml:"address"`
	Username  string                 `toml:"username"`
	Password  string                 `toml:"password"`
	Variables map[string]interface{} `toml:"variables,omitempty"`
}

type Group struct {
	Members   []string               `toml:"members"`
	Variables map[string]interface{} `toml:"variables"`
}

type Inventory struct {
	Machines map[string]Machine `toml:"machines"`
	Groups   map[string]Group   `toml:"groups"`
}

func (i Inventory) ResolveHosts(name string) []string {
	if machine, ok := i.Machines[name]; ok {
		return []string{machine.Address}
	} else if group, ok := i.Groups[name]; ok {
		var result []string
		for _, member := range group.Members {
			resolved := i.ResolveHosts(member)
			result = append(result, resolved...)
		}
		return result
	} else {
		return []string{}
	}
}
