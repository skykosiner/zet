package grep

import (
	"bufio"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/skykosiner/zet/pkg/color"
	"github.com/skykosiner/zet/pkg/config"
)

type match struct {
	line       string
	lineNumber string
	path       string
}

func (m match) String() string {
	return fmt.Sprintf("%s\n%s:%s\n", m.path, color.GreenString(m.lineNumber), m.line)
}

func NewMatch(line, lineNumber, path string) match {
	return match{
		line:       line,
		lineNumber: lineNumber,
		path:       path,
	}
}

func Grep(c config.Config, search string) {
	var matches []match
	searchRegex := regexp.MustCompile(search)
	err := filepath.Walk(c.Vault, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			defer file.Close()

			scanner := bufio.NewScanner(file)
			lineNum := 1
			for scanner.Scan() {
				line := scanner.Text()
				match := searchRegex.FindStringIndex(line)
				if len(match) > 0 {
					coloredLine := line[:match[0]] + color.RedString(line[match[0]:match[1]]) + line[match[1]:]
					path = strings.Replace(path, c.Vault+"/", "", -1)
					matches = append(matches, NewMatch(coloredLine, strconv.Itoa(lineNum), color.PurpleString(path)))
				}

				lineNum++
			}
		}

		return nil
	})

	if err != nil {
		slog.Error("Error when walking files", "error", err)
		return
	}

	for _, match := range matches {
		fmt.Println(match)
	}
}
