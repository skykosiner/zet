package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/skykosiner/zet/pkg/config"
	"github.com/skykosiner/zet/pkg/grep"
	"github.com/skykosiner/zet/pkg/notes"
	"github.com/skykosiner/zet/pkg/tags"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Short: "zet - Terminal Zettelkasten Manager",
		Use:   "zet",
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
		{
			Use:   "tags",
			Short: "Search for tag",
			Run: func(cmd *cobra.Command, args []string) {
				tags.Tags(c, fzfOptions)
			},
		},
		{
			Use:   "delete",
			Short: "Search for anote in your vault and delete it.",
			Run: func(cmd *cobra.Command, args []string) {
				notes.DeleteNote(c, fzfOptions)
			},
		},
		{
			Use:   "grep",
			Short: "Search notes with regex patterns",
			Run: func(cmd *cobra.Command, args []string) {
				grep.Grep(c, args[0])
			},
		},
		{
			Use:     "new",
			Short:   "Create a new note",
			Example: "zet new 'hello world'\nzet new 'python is cool' --path path/in/vault/for/new/note",
			Run: func(cmd *cobra.Command, args []string) {
				if len(args) < 1 {
					fmt.Println("Error: missing note name.")
					os.Exit(1)
				}

				notePath := fmt.Sprintf("%s/%s", c.Vault, newNoteFolder)
				notes.NewNote(args[0], notePath)
			},
		},
	}

	for _, command := range commands {
		rootCmd.AddCommand(&command)
	}

	if err := rootCmd.Execute(); err != nil {
		slog.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}
