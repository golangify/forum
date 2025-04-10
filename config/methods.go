package config

import (
	"encoding/json"
	"os"
)

const perm = 0666

func LoadFromJsonFile(path string) (*Config, error) {
	cfg := &Config{path: path}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = cfg.Save()
			if err != nil {
				return nil, err
			}
			return cfg, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	cfg.Save()

	return cfg, nil
}

func (c *Config) Save() error {
	c.safe()
	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(c.path, data, perm)
}
