package main

import (
	"flag"
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
		- tommorow
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

	notePath := flag.String("path", c.NewNotePath, "Path to put the new note in.")
	flag.Parse()

	arg := os.Args[1]
	switch arg {
	default:
		notes.NewNote(os.Args[1], *notePath)
	}
}
