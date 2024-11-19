# Zet
* A tool for managing notes in [Obsidian](https://obsidian.md) from your terminal
* In order to use a lot of the searching features you're going to want to want `fzf`
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
* `daily_note.template`: The name of the template to use when craeting a new daily note
* `daily_note.daily_note_date_format`: The format of the date used in your daily notes
## Usage
* Any command that uses fzf can accept the optinal flag of `--fzf-options` such as:
    * `zet search --fzf-options "--preview='bat --color=always --style=numbers {}' --preview-window=bottom:80%"`
* `zet today`
    * Open today's daily note
* `zet new-entry`
    * Open the current daily note and append the current time as a level two header
* `zet yesterday`
    * Open yesterdays daily note
* `zet tomorrow`
    * Open tomorrow's daily note
* `zet daily`
    * Search over all your daily notes and pick one to open
* `zet new <name of note>`
    * The note will be created where you've told the config to create the new note
    * `zet new <name of note> --path <sub path of vault to put the new note in>`
* `zet search`
    * Search for a note using fzf
    * `zet search --folder <sub path in your vault to search under>`
* `zet delete`
    * Search for notes and the notes you press enter on will be deleted
## Get completion in your shell
### ZSH
```zsh
eval "$(zet completion zsh)"
```
### Bash
```bash
eval "$(zet completion bash)"
```
### Fish
* Not too sure how fish works lol, but you can use use a `zet completion fish`
too generate the stuff needed
