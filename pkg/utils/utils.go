package utils

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
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

func SearchFZF(fzfOptions string, data []string) string {
	command := fmt.Sprintf("echo -e \"%s\" | fzf %s", strings.Join(data, "\\n"), fzfOptions)
	cmd := exec.Command("bash", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		fmt.Printf("Output: %s\n", string(output))
		panic(err)
	}

	return strings.TrimSpace(string(output))
}
