package notes

import (
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"strings"

	"github.com/skykosiner/zet/pkg/config"
	"github.com/skykosiner/zet/pkg/utils"
)

func getFiles(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && !strings.Contains(path, ".git") && !strings.Contains(path, ".obsidian") && strings.HasSuffix(info.Name(), ".md") {
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
	utils.OpenInEditor(utils.SearchFZF(fzfOptions, fmt.Sprintf("echo -e \"%s\"", strings.Join(files, "\n"))))
}
