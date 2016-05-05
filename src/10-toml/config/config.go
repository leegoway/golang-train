package config

type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
}

type song struct {
	Author string
	Name   string
}

type redis struct {
	Host    string
	Ports   []int
	Enabled bool
}

type Config struct {
	Title string
	Owner ownerInfo
	Redis redis `toml:"redis"`
	Songs map[string]song
}
