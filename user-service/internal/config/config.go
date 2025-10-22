package config

type Config struct {
	DB struct {
	}
}

func NewConfig() *Config {
	return &Config{}
}
