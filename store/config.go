package store

type Config struct {
	ConnStr          string `toml:"connection_string"`
	MigrationsFolder string `toml:"migrations_folder"`
}

func NewConfig() *Config {
	return &Config{
		ConnStr:          "localhost:5432",
		MigrationsFolder: "./migrations",
	}
}
