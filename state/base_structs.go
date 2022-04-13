package state

type State struct {
	Hosts   string   `toml:"hosts"`
	Scripts []string `toml:"scripts"`
}

type Set map[string]State
