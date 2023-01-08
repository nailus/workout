package apiserver

import (
	"github.com/nailus/workout/internal/database"
)

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	Database *database.Config
}

func NewConfig() *Config {
	return &Config{
		Database: database.NewConfig(),
	}
}