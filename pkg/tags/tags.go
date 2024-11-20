package tags

import (
	"bufio"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/skykosiner/zet/pkg/config"
	"github.com/skykosiner/zet/pkg/utils"
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

func getTags(c config.Config, tagRegex *regexp.Regexp) []string {
	var tags []string
	tagSet := make(map[string]struct{})
	// TODO: Get my regex license!

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
		return tags
	}

	for tag := range tagSet {
		tags = append(tags, tag)
	}

	sort.Strings(tags)
	return tags
}

// TODO: add in yaml stuff
func Tags(c config.Config, fzfOptions string) {
	tags := getTags(c, regexp.MustCompile(`#(\w\S*)`))

	if len(tags) == 0 {
		slog.Info("Can't fand any tags.")
		return
	}

	tag := utils.SearchFZF(fzfOptions, fmt.Sprintf("echo -e '%s'", strings.Join(tags, "\n")))
	fmt.Println(tag)
}

func SearchByTag(c config.Config, fzfOptions, tag string) {
	if len(tag) == 0 {
		slog.Error("Please provide a tag.")
		return
	}

}
