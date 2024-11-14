package utils

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
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

func GetFiles(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}
