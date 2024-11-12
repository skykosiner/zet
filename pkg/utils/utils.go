package utils

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

func OpenInEditor(path string) {
	editor := os.Getenv("EDITOR")
	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		slog.Error("Error opening up new note, but it has been set.", "error", err)
		os.Exit(0)
	}
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CatOrBat() string {
	_, err := exec.LookPath("bat")

	if err == nil {
		fmt.Println("test")
		return "bat --color=always --style=numbers"
	}

	return "cat"
}
