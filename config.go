package main

import "runtime"

type CmdEntry struct {
	Command  string
	Platform string
}

type Config struct {
	View map[string]map[string][]CmdEntry
	Open map[string]map[string][]CmdEntry
	Exec map[string]map[string][]CmdEntry
}

func (c *Config) getCmds(action, matchType, key string) []string {
	var m map[string]map[string][]CmdEntry
	switch action {
	case "view":
		m = c.View
	case "open":
		m = c.Open
	case "exec":
		m = c.Exec
	}
	if m != nil {
		for _, entry := range m[matchType][key] {
			if entry.Platform == "" || entry.Platform == runtime.GOOS {
				return []string{entry.Command}
			}
		}
	}
	return nil
}

var DefaultConfig = &Config{
	View: map[string]map[string][]CmdEntry{
		"ext": {
			"default":    {{Command: "cat"}},
			"md":         {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"txt":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"log":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"json":       {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"yaml":       {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"yml":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"diff":       {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"go":         {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"rs":         {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"py":         {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"sh":         {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"bash":       {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"zsh":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"ts":         {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"tsx":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"js":         {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"jsx":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"vue":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"svelte":     {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"c":          {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"cpp":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"h":          {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"hpp":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"lua":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"vim":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"toml":       {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"xml":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"html":       {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"css":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"scss":       {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"sass":       {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"less":       {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"sql":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"graphql":    {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"mdx":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"jsonc":      {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"ini":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"env":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"conf":       {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"cfg":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"properties": {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"java":       {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"kt":         {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"swift":      {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"rb":         {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"php":        {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"cs":         {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"tar":        {{Command: "atool --list"}},
			"rar":        {{Command: "atool --list"}},
			"zip":        {{Command: "atool --list"}},
			"7z":         {{Command: "atool --list"}},
			"gz":         {{Command: "atool --list"}},
			"bz2":        {{Command: "atool --list"}},
			"xz":         {{Command: "atool --list"}},
			"tgz":        {{Command: "atool --list"}},
		},
		"mimetype": {
			"default":            {{Command: "cat"}},
			"text/plain":         {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"text/x-shellscript": {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
			"text/diff":          {{Command: "bat --color=always --paging=never --tabs=2 --line-range :500"}},
		},
	},
	Open: map[string]map[string][]CmdEntry{
		"ext": {
			"default": {
				{Command: "xdg-open"},
				{Command: "open", Platform: "darwin"},
			},
			"md":         {{Command: "nvim"}},
			"txt":        {{Command: "nvim"}},
			"log":        {{Command: "nvim"}},
			"json":       {{Command: "nvim"}},
			"yaml":       {{Command: "nvim"}},
			"yml":        {{Command: "nvim"}},
			"diff":       {{Command: "nvim"}},
			"go":         {{Command: "nvim"}},
			"rs":         {{Command: "nvim"}},
			"py":         {{Command: "nvim"}},
			"sh":         {{Command: "nvim"}},
			"bash":       {{Command: "nvim"}},
			"zsh":        {{Command: "nvim"}},
			"ts":         {{Command: "nvim"}},
			"tsx":        {{Command: "nvim"}},
			"js":         {{Command: "nvim"}},
			"jsx":        {{Command: "nvim"}},
			"vue":        {{Command: "nvim"}},
			"svelte":     {{Command: "nvim"}},
			"c":          {{Command: "nvim"}},
			"cpp":        {{Command: "nvim"}},
			"h":          {{Command: "nvim"}},
			"hpp":        {{Command: "nvim"}},
			"lua":        {{Command: "nvim"}},
			"vim":        {{Command: "nvim"}},
			"toml":       {{Command: "nvim"}},
			"xml":        {{Command: "nvim"}},
			"html":       {{Command: "nvim"}},
			"css":        {{Command: "nvim"}},
			"scss":       {{Command: "nvim"}},
			"sass":       {{Command: "nvim"}},
			"less":       {{Command: "nvim"}},
			"sql":        {{Command: "nvim"}},
			"graphql":    {{Command: "nvim"}},
			"mdx":        {{Command: "nvim"}},
			"jsonc":      {{Command: "nvim"}},
			"ini":        {{Command: "nvim"}},
			"env":        {{Command: "nvim"}},
			"conf":       {{Command: "nvim"}},
			"cfg":        {{Command: "nvim"}},
			"properties": {{Command: "nvim"}},
			"java":       {{Command: "nvim"}},
			"kt":         {{Command: "nvim"}},
			"swift":      {{Command: "nvim"}},
			"rb":         {{Command: "nvim"}},
			"php":        {{Command: "nvim"}},
			"cs":         {{Command: "nvim"}},
			"pdf": {
				{Command: "zathura"},
				{Command: "open", Platform: "darwin"},
			},
			"png": {
				{Command: "feh"},
				{Command: "open", Platform: "darwin"},
			},
			"jpg": {
				{Command: "sxiv"},
				{Command: "open", Platform: "darwin"},
			},
			"jpeg": {
				{Command: "sxiv"},
				{Command: "open", Platform: "darwin"},
			},
			"gif": {
				{Command: "sxiv -a"},
				{Command: "open", Platform: "darwin"},
			},
			"svg": {
				{Command: "sxiv -a"},
				{Command: "open", Platform: "darwin"},
			},
			"webp": {
				{Command: "sxiv"},
				{Command: "open", Platform: "darwin"},
			},
			"bmp": {
				{Command: "feh"},
				{Command: "open", Platform: "darwin"},
			},
			"tiff": {
				{Command: "feh"},
				{Command: "open", Platform: "darwin"},
			},
			"ico": {
				{Command: "feh"},
				{Command: "open", Platform: "darwin"},
			},
			"mp3": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"mp4": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"mov": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"avi": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"mkv": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"webm": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"flv": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"wmv": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"m4v": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"wav": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"flac": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"ogg": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"m4a": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"aac": {
				{Command: "mpv"},
				{Command: "open", Platform: "darwin"},
			},
			"ttf": {
				{Command: "xdg-open"},
				{Command: "open", Platform: "darwin"},
			},
			"otf": {
				{Command: "xdg-open"},
				{Command: "open", Platform: "darwin"},
			},
			"woff": {
				{Command: "xdg-open"},
				{Command: "open", Platform: "darwin"},
			},
			"woff2": {
				{Command: "xdg-open"},
				{Command: "open", Platform: "darwin"},
			},
			"tar": {
				{Command: "atool --list"},
				{Command: "tar -tf", Platform: "darwin"},
			},
			"rar": {
				{Command: "atool --list"},
				{Command: "open", Platform: "darwin"},
			},
			"zip": {
				{Command: "atool --list"},
				{Command: "unzip -l", Platform: "darwin"},
			},
			"7z": {
				{Command: "atool --list"},
				{Command: "open", Platform: "darwin"},
			},
			"gz": {
				{Command: "atool --list"},
				{Command: "open", Platform: "darwin"},
			},
			"bz2": {
				{Command: "atool --list"},
				{Command: "open", Platform: "darwin"},
			},
			"xz": {
				{Command: "atool --list"},
				{Command: "open", Platform: "darwin"},
			},
			"tgz": {
				{Command: "atool --list"},
				{Command: "open", Platform: "darwin"},
			},
		},
		"mimetype": {
			"default": {
				{Command: "xdg-open"},
				{Command: "open", Platform: "darwin"},
			},
			"text/diff":          {{Command: "nvim"}},
			"text/plain":         {{Command: "nvim"}},
			"text/x-shellscript": {{Command: "nvim"}},
		},
	},
	Exec: map[string]map[string][]CmdEntry{
		"ext": {
			"default": {{Command: "bash"}},
			"sh":      {{Command: "bash"}},
			"py":      {{Command: "python3"}},
			"go":      {{Command: "go run"}},
			"rs":      {{Command: "cargo run"}},
			"js":      {{Command: "node"}},
			"ts":      {{Command: "ts-node"}},
			"lua":     {{Command: "lua"}},
		},
		"mimetype": {},
	},
}
