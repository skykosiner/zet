package notes

import (
	"fmt"
	"log/slog"
	"os/exec"
	"strings"

	"github.com/skykosiner/zet/pkg/config"
	"github.com/skykosiner/zet/pkg/utils"
)

func SearchNotes(c config.Config, folder, fzfOptions string) {
	path := fmt.Sprintf("%s/%s", c.Vault, folder)
	files, err := utils.GetFiles(path)
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
