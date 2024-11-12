package notes

import (
	"fmt"
	"time"

	"github.com/skykosiner/zet/pkg/config"
	"github.com/skykosiner/zet/pkg/utils"
)

var currentDate = time.Now().Format("2006-01-02")

/*
- If the note doesn't exist creat it
  - Maybe add an option where the use can generate the note from a template

- If the note does exist maybe add the current time to the bottom of the note
with a ## so that the user can journal or whatever, or maybe make that optional?
*/
func TodayNote(c config.Config) {
	path := fmt.Sprintf("%s/%s/%s.md", c.Vault, c.DailyNotes, currentDate)
	utils.OpenInEditor(path)
}

func NewEntry() {
}
