package cfg

import "fmt"

type DBConfig struct {
	User     string
	Password string
	HostName string
	Port     int
	Database string
}

func (config DBConfig) GetConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
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
			User:     "james",
			Password: "",
			HostName: "localhost",
			Port:     5432,
			Database: "personal_site",
		},
	}
}
