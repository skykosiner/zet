package tags

import (
	"bufio"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/skykosiner/zet/pkg/config"
)

func exractTags(path string, tagSet map[string]struct{}, tagRegex *regexp.Regexp) {
	file, err := os.Open(path)
	if err != nil {
		slog.Error("Error opening file.", "error", err, "file path", path)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := tagRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 1 {
				tagSet[match[1]] = struct{}{}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		slog.Error("Error reading file.", "error", err, "file path", path)
		return
	}
}

// TODO: add in yaml stuff
func Tags(c config.Config) {
	tagSet := make(map[string]struct{})
	tagRegex := regexp.MustCompile(`#(\w+[\w-]*)`)

	err := filepath.Walk(c.Vault, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			exractTags(path, tagSet, tagRegex)
		}

		return nil
	})

	if err != nil {
		slog.Error("Error getting tags.", "error", err)
		return
	}

	var tags []string
	for tag := range tagSet {
		tags = append(tags, tag)
	}

	sort.Strings(tags)
}
