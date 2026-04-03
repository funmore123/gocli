package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config 存储认证凭证
type Config struct {
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

func Dir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".gocli")
}

func Path() string {
	return filepath.Join(Dir(), "config.json")
}

func Load() (*Config, error) {
	data, err := os.ReadFile(Path())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("not logged in, run: gocli auth login --api-key <key> --api-secret <secret>")
		}
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func Remove() error {
	return os.Remove(Path())
}

func Save(cfg *Config) error {
	if err := os.MkdirAll(Dir(), 0700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(Path(), data, 0600)
}
