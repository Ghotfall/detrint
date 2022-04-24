package inv

type Group struct {
	Members   []string               `toml:"members" json:"members"`
	Variables map[string]interface{} `toml:"variables,omitempty" json:"variables,omitempty"`
}
