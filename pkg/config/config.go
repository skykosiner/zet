package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type DailyNote struct {
	Template            string `json:"template"`
	DailyNotes          string `json:"daily_notes"`
	DailyNoteDateFormat string `json:"daily_note_date_format"`
}

type Config struct {
	Vault         string    `json:"vault"`
	TemplatesPath string    `json:"templates_path"`
	NewNotePath   string    `json:"new_note_path"`
	DailyNote     DailyNote `json:"daily_note"`
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
