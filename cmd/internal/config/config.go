package config

import (
	"fmt"
	"path/filepath"

	"github.com/kndrad/squil/cmd/internal/logging"
	"github.com/kndrad/squil/internal/shelter"
	"github.com/spf13/viper"
)

const DefaultEnvFilePath = ".env"

func LoadShelterConfig(path string) (*shelter.Config, error) {
	logger := logging.DefaultLogger()
	logger.Info("Loading shelter database configuration")

	if path == "" {
		path = DefaultEnvFilePath
	}
	viper.SetConfigFile(filepath.Clean(path))

	if err := viper.ReadInConfig(); err != nil {
		if _, notfound := err.(viper.ConfigFileNotFoundError); notfound {
			return nil, fmt.Errorf("config file not found: %w", err)
		} else {
			return nil, fmt.Errorf("reading in config: %w", err)
		}
	}

	viper.AutomaticEnv()

	cfg := &shelter.Config{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		DBName:   viper.GetString("DB_NAME"),
	}
	logger.Info("Config", "db_host", cfg.Host)

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return cfg, nil
}
