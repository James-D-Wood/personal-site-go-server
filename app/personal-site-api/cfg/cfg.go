package cfg

import (
	"fmt"

	"github.com/spf13/viper"
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
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return &Config{
		Database: DBConfig{
			User:     viper.GetString("postgresql.user"),
			Password: viper.GetString("postgresql.password"),
			HostName: viper.GetString("postgresql.host"),
			Port:     viper.GetString("postgresql.port"),
			Database: viper.GetString("postgresql.database"),
		},
	}
}
