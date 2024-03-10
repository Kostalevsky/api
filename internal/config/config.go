package config

type Config struct {
	Transport Transport `yaml:"transport"`
	Repo      Repo      `yaml:"repo"`
}
