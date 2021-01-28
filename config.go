package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFile = ".huetility.json"

const configEnv = "HUETILITY_CONFIG"

type Config struct {
	IPAddress string `json:"ip_address"`
	UserName  string `json:"user_name"`
	Scenes    Scenes
}

type Scenes []*Scene

func (ss Scenes) get(id string) (*Scene, error) {
	for _, s := range ss {
		if s.ID == id {
			return s, nil
		}
	}

	return nil, ErrIDNotFound(id)
}

type Scene struct {
	ID      string     `json:"id"`
	LightID HueLightID `json:"light_id"`
	Name    string     `json:"name"`
	State   *HueState  `json:"state"`
}

func loadConfig() (*Config, error) {
	configFilePath := os.Getenv(configEnv)
	if configFilePath == "" {
		var err error

		configFilePath, err = defaultConfigFilePath()
		if err != nil {
			return nil, err
		}
	}

	f, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := json.NewDecoder(f).Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func defaultConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, configFile), nil
}
