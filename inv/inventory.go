package inv

import "fmt"

type Inventory struct {
	Machines map[string]Machine `toml:"machines" json:"machines"`
	Groups   map[string]Group   `toml:"groups" json:"groups"`
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

func (i Inventory) GetMachineVars(servername string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	machine, ok := i.Machines[servername]
	if !ok {
		// return empty map, not null
		return result, fmt.Errorf("can't find machine with name '%s'", servername)
	}

	for _, group := range i.Groups {
		for _, member := range group.Members {
			if member == servername {
				for k, v := range group.Variables {
					result[k] = v
				}
			}
		}
	}

	for k, v := range machine.Variables {
		result[k] = v
	}

	return result, nil
}
