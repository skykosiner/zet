package notes

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/skykosiner/zet/pkg/config"
	"github.com/skykosiner/zet/pkg/utils"
)

func DeleteNote(c config.Config, fzfOptions string) {
	files, err := getFiles(c.Vault)
	if err != nil {
		slog.Error("Error getting files in your vault.", "error", err)
		return
	}

	path := utils.SearchFZF(fzfOptions, fmt.Sprintf("echo -e \"%s\"", strings.Join(files, "\n")))
	if err := os.Remove(path); err != nil {
		slog.Error("Error deleting note", "error", err, "path", path)
		return
	}
}
