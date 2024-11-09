package main

import (
	"log/slog"
	"os"

	"github.com/skykosiner/zet/pkg/config"
)

func main() {
	/*
		options:
		- default arg will juts become a new note in _Inbox folder
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

	arg := os.Args[1]
	switch arg {
	default:
	}
}
