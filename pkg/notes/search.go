package notes

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/skykosiner/zet/pkg/config"
)

// Walk a path provided and get all file names in the root and sub dirs
func getFiles(path string) ([]string, error) {
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

func SearchNotes(c config.Config, folder, fzfOptions string) {
	path := fmt.Sprintf("%s/%s", c.Vault, folder)
	files, err := getFiles(path)
	if err != nil {
		slog.Error("Erorr getting files.", "error", err, "path", path)
		return
	}

	command := fmt.Sprintf("printf '%s' | fzf %s", strings.Join(files, "\n"), fzfOptions)
	cmd := exec.Command("bash", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		fmt.Printf("Output: %s\n", string(output))
		panic(err)
	}

	fmt.Println(string(output))
}
