package inv

type Group struct {
	Members   []string               `toml:"members"`
	Variables map[string]interface{} `toml:"variables"`
}
