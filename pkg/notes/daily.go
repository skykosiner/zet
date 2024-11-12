package notes

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/skykosiner/zet/pkg/config"
	"github.com/skykosiner/zet/pkg/utils"
)

func insertDailyTemplate(c config.Config) string {
	currentDate := time.Now().Format(c.DailyNote.DailyNoteDateFormat)
	path := fmt.Sprintf("%s/%s/%s.md", c.Vault, c.DailyNote.DailyNotes, currentDate)

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
	path := insertDailyTemplate(c)
	utils.OpenInEditor(path)
}

func NewEntry(c config.Config) {
	path := insertDailyTemplate(c)
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
