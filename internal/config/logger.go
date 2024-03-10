package config

type Logger struct {
	LogFilePath string `yaml:"path"`
	Level       string `yaml:"level"`
	MaxSize     int    `yaml:"maxSizeMBytes"`
	MaxAge      int    `yaml:"maxAgeHours"`
	MaxBackups  int    `yaml:"maxBackups"`
	Compress    bool   `yaml:"compress"`
}
