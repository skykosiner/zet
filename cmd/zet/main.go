package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/skykosiner/zet/pkg/config"
	"github.com/skykosiner/zet/pkg/notes"
)

func main() {
	/*
		options:
		- default arg will juts become a new note in the path a user has chose in the config
			- --path chose the path to create the new note in
		- today
		- yesterday
		- tomorrow
		- daily
			- Pick any daily note date
		- tags
		- tag
			- Search by tag
	*/

	c, err := config.NewConfig()
	if err != nil {
		slog.Error("Error with config. Please check the README.", "error", err)
		os.Exit(1)
	}

	newNoteFolder := flag.String("path", c.NewNotePath, "Path to put the new note in.")
	flag.Parse()
	notePath := fmt.Sprintf("%s/%s", c.Vault, *newNoteFolder)
	args := flag.Args()

	switch args[0] {
	case "today":
		notes.TodayNote(c)
	case "yesterday":
		notes.YesterdaysNote(c)
	case "tomorrow":
		notes.TomorrowsNote(c)
	case "new-entry":
		notes.NewEntry(c)
	default:
		notes.NewNote(args[0], notePath)
	}
}
