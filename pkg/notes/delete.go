package notes

import (
	"log/slog"
	"os"

	"github.com/skykosiner/zet/pkg/config"
	"github.com/skykosiner/zet/pkg/utils"
)

func DeleteNote(c config.Config, fzfOptions string) {
	files, err := getFiles(c.Vault)
	if err != nil {
		slog.Error("Error getting files in your vault.", "error", err)
		return
	}

	path := utils.SearchFZF(fzfOptions, files)
	if err := os.Remove(path); err != nil {
		slog.Error("Error deleting note", "error", err, "path", path)
		return
	}
}
