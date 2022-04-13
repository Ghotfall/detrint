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
