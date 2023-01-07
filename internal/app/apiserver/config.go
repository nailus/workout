package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":2000",
		//LogLevel: "debug",
	}
}