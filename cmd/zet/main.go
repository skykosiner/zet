package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/skykosiner/zet/pkg/config"
	"github.com/skykosiner/zet/pkg/notes"
	"github.com/spf13/cobra"
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
	rootCmd := &cobra.Command{
		Short:   "zet - Terminal Zettelkasten Manager",
		Use:     "zet",
		Example: "zet [new note name]",
	}

	c, err := config.NewConfig()
	if err != nil {
		slog.Error("Error with config. Please check the README.", "error", err)
		os.Exit(1)
	}

	// Global flag for specifying the path
	var newNoteFolder string
	var fzfOptions string
	var folder string
	rootCmd.PersistentFlags().StringVar(&newNoteFolder, "path", c.NewNotePath, "Path to put the new note in.")
	rootCmd.PersistentFlags().StringVar(&fzfOptions, "fzf-options", "", "Additional options to pass to fzf.\nzet search --fzf-options \"--preview='bat --color=always --style=numbers {}' --preview-window='bottom,90%'\"")
	rootCmd.PersistentFlags().StringVar(&folder, "folder", "", "Which sub folder/folders to search when looking for notes.")

	commands := []cobra.Command{
		{
			Use:   "today",
			Short: "Show today's note",
			Run: func(cmd *cobra.Command, args []string) {
				notes.TodayNote(c)
			},
		},
		{
			Use:   "yesterday",
			Short: "Show yesterday's note",
			Run: func(cmd *cobra.Command, args []string) {
				notes.YesterdaysNote(c)
			},
		},
		{
			Use:   "tomorrow",
			Short: "Show tomorrow's note",
			Run: func(cmd *cobra.Command, args []string) {
				notes.TomorrowsNote(c)
			},
		},
		{
			Use:   "daily",
			Short: "Pick a daily note date",
			Run: func(cmd *cobra.Command, args []string) {
				notes.SelectDaily(c, fzfOptions)
			},
		},
		{
			Use:   "new-entry",
			Short: "Create a new entry",
			Run: func(cmd *cobra.Command, args []string) {
				notes.NewEntry(c)
			},
		},
		{
			Use:     "search",
			Short:   "Search notes",
			Example: "zet search\nzet search --folder sub_folder/another_folder",
			Run: func(cmd *cobra.Command, args []string) {
				notes.SearchNotes(c, folder, fzfOptions)
			},
		},
	}

	for _, command := range commands {
		rootCmd.AddCommand(&command)
	}

	// Default behavior for creating a new note
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: missing note name.")
			os.Exit(1)
		}

		notePath := fmt.Sprintf("%s/%s", c.Vault, newNoteFolder)
		notes.NewNote(args[0], notePath)
	}

	rootCmd.Example = "zet 'new note name'\nzet 'new note name' --path sub_path_in_vault/path_two"

	if err := rootCmd.Execute(); err != nil {
		slog.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}
