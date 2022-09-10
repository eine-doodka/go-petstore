package apiserver

type Config struct {
	BindAddr           string `toml:"bind_addr"`
	LogLevel           string `toml:"log_level"`
	DbConnectionString string `toml:"db_connection_string"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: "localhost:8080",
		LogLevel: "debug",
	}
}
