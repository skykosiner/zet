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
		Use: "zet",
	}

	c, err := config.NewConfig()
	if err != nil {
		slog.Error("Error with config. Please check the README.", "error", err)
		os.Exit(1)
	}

	// Global flag for specifying the path
	var newNoteFolder string
	var fzfOptions string
	rootCmd.PersistentFlags().StringVar(&newNoteFolder, "path", c.NewNotePath, "Path to put the new note in.")
	rootCmd.PersistentFlags().StringVar(&fzfOptions, "fzf-options", "", "Additional options to pass to fzf.")

	// Subcommand for "today"
	todayCmd := &cobra.Command{
		Use:   "today",
		Short: "Show today's note",
		Run: func(cmd *cobra.Command, args []string) {
			notes.TodayNote(c)
		},
	}

	// Subcommand for "yesterday"
	yesterdayCmd := &cobra.Command{
		Use:   "yesterday",
		Short: "Show yesterday's note",
		Run: func(cmd *cobra.Command, args []string) {
			notes.YesterdaysNote(c)
		},
	}

	// Subcommand for "tomorrow"
	tomorrowCmd := &cobra.Command{
		Use:   "tomorrow",
		Short: "Show tomorrow's note",
		Run: func(cmd *cobra.Command, args []string) {
			notes.TomorrowsNote(c)
		},
	}

	// Subcommand for "daily"
	dailyCmd := &cobra.Command{
		Use:   "daily",
		Short: "Pick a daily note date",
		Run: func(cmd *cobra.Command, args []string) {
			notes.SelectDaily(c, fzfOptions)
		},
	}

	// Subcommand for "new-entry"
	newEntryCmd := &cobra.Command{
		Use:   "new-entry",
		Short: "Create a new entry",
		Run: func(cmd *cobra.Command, args []string) {
			notes.NewEntry(c)
		},
	}

	// Default behavior for creating a new note
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: missing note name.")
			os.Exit(1)
		}

		// Combine the path from config and the user-specified folder
		notePath := fmt.Sprintf("%s/%s", c.Vault, newNoteFolder)
		notes.NewNote(args[0], notePath)
	}

	// Add subcommands to the root command
	rootCmd.AddCommand(todayCmd, yesterdayCmd, tomorrowCmd, dailyCmd, newEntryCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}
