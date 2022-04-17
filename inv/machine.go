package inv

type Machine struct {
	Address   string                 `toml:"address"`
	Username  string                 `toml:"username"`
	Password  string                 `toml:"password"`
	Variables map[string]interface{} `toml:"variables,omitempty"`
}
