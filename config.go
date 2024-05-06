package main

import (
	"gopkg.in/yaml.v3"
	"log/slog"
	"os"
)

type Config struct {
	Environment string `yaml:"environment"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
}

func NewDefaultConfig() *Config {
	return &Config{
		Environment: "development",
		Host:        "localhost",
		Port:        "8080",
	}
}

// GetConfig reads the config file from the given path and returns a Config struct, if it does not exist or is invalid it will return a default Config struct
func GetConfig(path string) (*Config, error) {
	slog.Info("Reading config file", slog.String("path", path))
	config := NewDefaultConfig()

	_, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		slog.Warn("Config file not found, using default config")

		configBytes, err := yaml.Marshal(config)
		if err != nil {
			slog.Error("Failed to marshal default config")
			return nil, err
		}

		err = os.WriteFile(path, configBytes, 0644)
		if err != nil {
			slog.Error("Failed to write default config to file")
			return nil, err
		}

		return config, nil
	}

	file, err := os.ReadFile(path)
	if err != nil {
		slog.Error("Failed to read config file", slog.String("path", path), slog.String("error", err.Error()))
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		slog.Error("Failed to unmarshal config file", slog.String("path", path), slog.String("error", err.Error()))
		return nil, err
	}

	return config, nil
}

// Save saves the config to the given path
func (c *Config) Save(path string) error {
	slog.Info("Saving config to file", slog.String("path", path))
	file, err := yaml.Marshal(c)
	if err != nil {
		slog.Error("Failed to marshal config", slog.String("error", err.Error()))
		return err
	}

	err = os.WriteFile(path, file, 0644)
	if err != nil {
		slog.Error("Failed to write config to file", slog.String("path", path), slog.String("error", err.Error()))
		return err
	}

	return nil
}
