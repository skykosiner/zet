package notes

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"

	"github.com/skykosiner/zet/pkg/config"
)

func NewNote(name string, c config.Config) {
	name = strings.Replace(name, " ", "-", 0)
	newNotePath := fmt.Sprintf("%s/%s.md", c.NewNotePath, name)
	editor := os.Getenv("EDITOR")

	if _, err := os.Create(newNotePath); err != nil {
		slog.Error("Erorr creating new note.", "error", err)
		os.Exit(0)
	}

	cmd := exec.Command(editor, newNotePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		slog.Error("Error opening up new note, but it has been set.", "error", err)
		os.Exit(0)
	}
}
