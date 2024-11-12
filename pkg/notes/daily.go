package notes

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/skykosiner/zet/pkg/config"
	"github.com/skykosiner/zet/pkg/utils"
)

func insertDailyTemplate(c config.Config, date string) string {
	path := fmt.Sprintf("%s/%s/%s.md", c.Vault, c.DailyNote.DailyNotes, date)

	if !utils.FileExists(path) {
		InsertTemplate(fmt.Sprintf("%s/%s/%s.md", c.Vault, c.TemplatesPath, c.DailyNote.Template), path)
	}

	return path
}

/*
- If the note doesn't exist creat it
  - Maybe add an option where the use can generate the note from a template

- If the note does exist maybe add the current time to the bottom of the note
with a ## so that the user can journal or whatever, or maybe make that optional?
*/
func TodayNote(c config.Config) {
	currentDate := time.Now().Format(c.DailyNote.DailyNoteDateFormat)
	path := insertDailyTemplate(c, currentDate)
	utils.OpenInEditor(path)
}

func TomorrowsNote(c config.Config) {
	tomorrowsDate := time.Now().AddDate(0, 0, +1).Format(c.DailyNote.DailyNoteDateFormat)
	path := insertDailyTemplate(c, tomorrowsDate)
	utils.OpenInEditor(path)
}

func YesterdaysNote(c config.Config) {
	yesterdaysDate := time.Now().AddDate(0, 0, -1).Format(c.DailyNote.DailyNoteDateFormat)
	path := fmt.Sprintf("%s/%s/%s.md", c.Vault, c.DailyNote.DailyNotes, yesterdaysDate)

	if !utils.FileExists(path) {
		slog.Info("You don't have a note from yesterday.", "Date", yesterdaysDate)
		return
	}

	utils.OpenInEditor(path)
}

func NewEntry(c config.Config) {
	currentDate := time.Now().Format(c.DailyNote.DailyNoteDateFormat)
	path := insertDailyTemplate(c, currentDate)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		slog.Error("Cloudn't open daily note file to append new entry.", "error", err, "path", path)
		return
	}

	defer f.Close()

	// TODO: Make the minutes format add in a double 00 for 0-9 minute and not
	// just do a :3 but instead :03
	if _, err = f.WriteString(fmt.Sprintf("## %d:%d\n", time.Now().Hour(), time.Now().Minute())); err != nil {
		panic(err)
	}

	utils.OpenInEditor(path)
}

func SelectDaily(c config.Config) {
	command := fmt.Sprintf(`ls "%s/%s" | fzf --preview="%s %s/%s/{}" --preview-window="80%%"`, c.Vault, c.DailyNote.DailyNotes, utils.CatOrBat(), c.Vault, c.DailyNote.DailyNotes)
	cmd := exec.Command("bash", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		fmt.Printf("Output: %s\n", string(output))
		panic(err)
	}

	path := fmt.Sprintf("%s/%s/%s", c.Vault, c.DailyNote.DailyNotes, strings.TrimSpace(string(output)))
	utils.OpenInEditor(path)
}
