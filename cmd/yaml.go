package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type yamlData struct {
	tags []string `yaml:"tags"`
}

func main() {
	file, _ := os.ReadFile("./test.md")
	var lines []string
	for _, line := range strings.Split(string(file), "\n") {
		lines = append(lines, line)
	}

	var yData yamlData
	var inYaml bool
	yamlStr := ""
	for idx, line := range lines {
		if idx == 0 && line == "---" {
			inYaml = true
		} else if inYaml && line == "---" {
			return
		} else {
			yamlStr += line
		}
	}

	fmt.Println("test", yamlStr)

	if err := yaml.Unmarshal([]byte(yamlStr), &yData); err != nil {
		slog.Error("Error converting yaml string to yaml.", "error", err)
		return
	}

	fmt.Println(yData)
}
