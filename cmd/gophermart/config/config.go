package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v10"
)

const (
	defaultServerAddress = ":8080"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	DBDSN         string `env:"DATABASE_DSN"`
	DevMode       bool   `env:"DEV_MODE"`
}

func (cfg *Config) String() string {
	return fmt.Sprintf(`
	ServerAddress: %s,
	DATABASE_DSN: "%s"
	`, cfg.ServerAddress, cfg.DBDSN)
}

func (cfg *Config) Load() {
	cfg.ServerAddress = defaultServerAddress
	cfg.DevMode = true
	cfg.parseFlags()
	cfg.parseEnv()
}

func (cfg *Config) parseEnv() {
	err := env.Parse(cfg)
	if err != nil {
		fmt.Println("Unable to load config:", err)
	}
}

func (cfg *Config) parseFlags() {
	flag.Func("a", "Example: -a localhost:8080", func(v string) error {
		cfg.ServerAddress = v
		return nil
	})

	cfg.DevMode = *flag.Bool("dev", false, "Example: -dev 1")

	flag.Func("d", "Example -d postgres://username:password@localhost:5432/database_name", func(v string) error {
		cfg.DBDSN = v
		return nil
	})
	flag.Parse()
}
