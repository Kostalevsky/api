package config

type Repo struct {
	User     string `yaml:"user"`
	Pass     string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"db"`
	SSLMode  string `yaml:"ssl"`
}
