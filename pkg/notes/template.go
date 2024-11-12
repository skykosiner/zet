package notes

import (
	"log/slog"
	"os"
)

func InsertTemplate(templatePath string, notePath string) {
	bytes, err := os.ReadFile(templatePath)
	if err != nil {
		slog.Error("Error reading template.", "error", err, "path", templatePath)
		os.Exit(1)
	}

	if err := os.WriteFile(notePath, bytes, 0644); err != nil {
		slog.Error("Error writing template.", "error", err, "path", templatePath, "noet to write", notePath)
		os.Exit(1)
	}
}
