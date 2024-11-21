package tags

import (
	"bufio"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/skykosiner/zet/pkg/config"
	"github.com/skykosiner/zet/pkg/utils"
	"gopkg.in/yaml.v3"
)

type tag struct {
	tag  string
	path string
}

type yamlData struct {
	Tags []string `yaml:"tags"`
}

type tags []tag

func (t tag) String() string {
	return fmt.Sprintf("%s: %s\n", t.path, t.tag)
}

func (t tags) getPaths() []string {
	var paths []string
	for _, tag := range t {
		paths = append(paths, tag.path)
	}

	return paths
}

func (t tags) getTags() []string {
	var tags []string
	for _, tag := range t {
		tags = append(tags, tag.tag)
	}

	return tags
}

func exractTags(path string, tagSet map[string]tag, tagRegex *regexp.Regexp) {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		slog.Error("Error opening file.", "error", err, "file path", path)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		matches := tagRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 1 {
				tagSet[match[1]] = tag{
					tag:  match[1],
					path: path,
				}
			}
		}
	}

	var yData yamlData
	var inYaml bool
	var yamlLines []string

	// Process lines
	for _, line := range lines {
		// Start YAML block
		if line == "---" {
			if inYaml {
				break // End of YAML block
			}
			inYaml = true
			continue
		}

		if inYaml {
			yamlLines = append(yamlLines, line)
		}
	}

	yamlStr := strings.Join(yamlLines, "\n")

	if err := yaml.Unmarshal([]byte(yamlStr), &yData); err != nil {
		slog.Error("Error converting yaml string to struct.", "error", err)
		return
	}

	for _, t := range yData.Tags {
		tagSet[t] = tag{
			tag:  t,
			path: path,
		}
	}

	if err := scanner.Err(); err != nil {
		slog.Error("Error reading file.", "error", err, "file path", path)
		return
	}
}

func getTags(c config.Config, tagRegex *regexp.Regexp) tags {
	var tags tags
	tagSet := make(map[string]tag)

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
		tags = append(tags, tagSet[tag])
	}

	return tags
}

func Tags(c config.Config, fzfOptions string) {
	// TODO: Get my regex license!
	tags := getTags(c, regexp.MustCompile(`#(\w\S*)`))
	if len(tags) == 0 {
		slog.Info("Can't fand any tags.")
		return
	}

	tag := utils.SearchFZF(fzfOptions, fmt.Sprintf("echo -e '%s'", strings.Join(tags.getTags(), "\n")))
	SearchByTag(c, fzfOptions, tag)
}

func SearchByTag(c config.Config, fzfOptions, tag string) {
	if len(tag) == 0 {
		slog.Error("Please provide a tag.")
		return
	}

	// TODO: Get my regex license!
	tags := getTags(c, regexp.MustCompile(fmt.Sprintf(`#(%s\S*)`, tag)))
	fmt.Println(tags)
	utils.OpenInEditor(utils.SearchFZF(fzfOptions, fmt.Sprintf("echo -e '%s'", strings.Join(tags.getPaths(), "\n"))))
}
