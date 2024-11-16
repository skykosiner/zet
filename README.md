# Zet
* A tool for managing notes in [Obsidian](https://obsidian.md) from your terminal
## Config
The config must be placed at `~/.config/zet/config.json`
```json
{
    "vault": "/home/sky/Documents/Linux-btw",
    "templates_path": "Templates",
    "new_note_path": "_Inbox",
    "daily_note": {
        "daily_notes": "Daily",
        "template": "Daily note",
        "daily_note_date_format": "2006-01-02"
    }
}
```
* `vault`: The path to your obsidian vault
* `templates_path`: The path to a folder where you keep your templates
* `new_note_path`: Which path in your vault to place a new note
* `daily_note.daily_notes`: The path to which folder your daily notes are stored in
* `daily_note.template`: The  name of the template to use when craeting a new daily note
* `daily_note.daily_note_date_format`:
## Usage
