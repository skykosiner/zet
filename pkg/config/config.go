package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Vault       string `json:"vault"`
	NewNotePath string `json:"new_note_path"`
}

func NewConfig() (Config, error) {
	var c Config
	bytes, err := os.ReadFile(fmt.Sprintf("%s/.config/zet/config.json", os.Getenv("HOME")))
	if err != nil {
		return c, err
	}

	if err := json.Unmarshal(bytes, &c); err != nil {
		return c, err
	}

	return c, nil
}
