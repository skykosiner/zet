package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type yamlData struct {
	Tags []string `yaml:"tags"`
}

func main() {
	// Read the file
	file, err := os.ReadFile("./test.md")
	if err != nil {
		slog.Error("Error reading file", "error", err)
		return
	}

	// Split lines
	lines := strings.Split(string(file), "\n")

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

	// Join YAML block into a string
	yamlStr := strings.Join(yamlLines, "\n")
	fmt.Println("YAML block:\n", yamlStr)

	// Unmarshal YAML
	if err := yaml.Unmarshal([]byte(yamlStr), &yData); err != nil {
		slog.Error("Error converting yaml string to struct.", "error", err)
		return
	}

	fmt.Println("Parsed YAML Data:", yData.Tags)
}
