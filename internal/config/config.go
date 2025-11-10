package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePAth() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return dir + string(os.PathSeparator) + configFileName, nil
}

func write(cfg Config) error {
	configPath, err := getConfigFilePAth()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Read() Config {
	configPath, err := getConfigFilePAth()
	if err != nil {
		return Config{}
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return Config{}
	}

	return config
}

func SetUser(userName string) error {
	cfg := Read()
	cfg.CurrentUserName = userName
	return write(cfg)
}
