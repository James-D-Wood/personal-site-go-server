package cfg

import (
	"fmt"
	"os"
)

type DBConfig struct {
	User     string
	Password string
	HostName string
	Port     string
	Database string
}

func (config DBConfig) GetConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.User,
		config.Password,
		config.HostName,
		config.Port,
		config.Database,
	)
}

type Config struct {
	Database DBConfig
}

func Load() *Config {
	return &Config{
		Database: DBConfig{
			User:     os.Getenv("PS_DB_User"),
			Password: os.Getenv("PS_DB_Password"),
			HostName: os.Getenv("PS_DB_HostName"),
			Port:     os.Getenv("PS_DB_Port"),
			Database: os.Getenv("PS_DB_Database"),
		},
	}
}
