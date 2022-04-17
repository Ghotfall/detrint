package inv

type Inventory struct {
	Machines map[string]Machine `toml:"machines"`
	Groups   map[string]Group   `toml:"groups"`
}

func (i Inventory) ResolveHosts(name string) map[string]Machine {
	if machine, ok := i.Machines[name]; ok {
		return map[string]Machine{name: machine}
	} else if group, ok := i.Groups[name]; ok {
		result := make(map[string]Machine)

		for _, member := range group.Members {
			resolved := i.ResolveHosts(member)
			for s, m := range resolved {
				result[s] = m
			}
		}

		return result
	} else {
		return map[string]Machine{}
	}
}
