package store

type Config struct {
	DatabaseURL	string `toml:"database_url"`
}

// initializing new default config
func NewConfig() *Config {
	return &Config{}
}

