package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

type Action func(*Lazy, *Config)

var Options = map[string]Action{
	"view": (*Lazy).VIEW,
	"open": (*Lazy).OPEN,
	"exec": (*Lazy).EXEC,
}

type Lazy struct {
	filePath string
	ext      string
	mimetype string
}

func NewLazy(filePath string) *Lazy {
	return &Lazy{
		filePath: filePath,
		ext:      getFileExt(filePath),
		mimetype: getFileMimeType(filePath),
	}
}

func (l *Lazy) VIEW(cfg *Config) {
	l.runCmd(cfg, "view")
}

func (l *Lazy) OPEN(cfg *Config) {
	l.runCmd(cfg, "open")
}

func (l *Lazy) EXEC(cfg *Config) {
	l.runCmd(cfg, "exec")
}

func (l *Lazy) runCmd(cfg *Config, action string) {
	var cmds []string

	switch action {
	case "view":
		cmds = append(cmds, cfg.View["ext"][l.ext]...)
		cmds = append(cmds, cfg.View["mimetype"][l.mimetype]...)
		cmds = append(cmds, cfg.View["ext"]["default"]...)
		cmds = append(cmds, cfg.View["mimetype"]["default"]...)
	case "open":
		cmds = append(cmds, cfg.Open["ext"][l.ext]...)
		cmds = append(cmds, cfg.Open["mimetype"][l.mimetype]...)
		cmds = append(cmds, cfg.Open["ext"]["default"]...)
		cmds = append(cmds, cfg.Open["mimetype"]["default"]...)
	case "exec":
		cmds = append(cmds, cfg.Exec["ext"][l.ext]...)
		cmds = append(cmds, cfg.Exec["mimetype"][l.mimetype]...)
		cmds = append(cmds, cfg.Exec["ext"]["default"]...)
		cmds = append(cmds, cfg.Exec["mimetype"]["default"]...)
	}

	for _, cmd := range cmds {
		finalCmd := fmt.Sprintf(`%s '%s'`, cmd, escapeShellArg(l.filePath))
		fmt.Printf("- cmd: %s\n\n", finalCmd)
		if err := l.exec(finalCmd); err == nil {
			return
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println("all commands failed")
}

func escapeShellArg(s string) string {
	s = strings.ReplaceAll(s, `'`, `'\\''`)
	return s
}

func (l *Lazy) exec(cmd string) error {
	c := exec.Command("bash", "-c", cmd)
	c.Stdout, c.Stderr = os.Stdout, os.Stderr
	return c.Run()
}

func getFileExt(filePath string) string {
	return strings.ToLower(strings.TrimPrefix(path.Ext(filePath), "."))
}

func getFileMimeType(filePath string) string {
	m, err := mimetype.DetectFile(filePath)
	if err != nil {
		return ""
	}
	parts := strings.SplitN(m.String(), ";", 2)
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func isFileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func (l *Lazy) PrintHelp() {
	fmt.Println(`
NAME
	lazy - A CLI tool that improves your work efficiency.

SYNOPSIS
	lazy -v
	lazy -h
	lazy -o <OPTION> -f <file>

DESCRIPTION
	lazy is a tool for CLI to view, open, execute files automatically.

OPTIONS
	-view    VIEW file with your default setting.
	-open    Open file with your default setting.
	-exec    Execute script with your default setting.

BUGS
	Report bugs to zetatez@icloud.com.
	`)
}

func (l *Lazy) PrintVersion() {
	fmt.Println("lazy version: 0.0.1")
}

func main() {
	help := flag.Bool("h", false, "help")
	version := flag.Bool("v", false, "version")
	option := flag.String("o", "", "operation (view, open, exec)")
	filePath := flag.String("f", "", "file path")
	flag.Parse()

	switch {
	case *help:
		NewLazy("").PrintHelp()
		return
	case *version:
		NewLazy("").PrintVersion()
		return
	}

	if *filePath == "" {
		fmt.Println("Error: file path is required. Use -h for help.")
		return
	}

	if !isFileExists(*filePath) {
		fmt.Println("Error: file does not exist.")
		return
	}

	cfg := DefaultConfig

	lazy := NewLazy(*filePath)
	fmt.Printf(`
- file path: %s
  - ext: %s
  - mimetype: %s
- option: %s
`,
		*filePath,
		lazy.ext,
		lazy.mimetype,
		*option,
	)
	if action, ok := Options[*option]; ok {
		action(lazy, cfg)
	} else {
		fmt.Println("Error: invalid option. Use -h for help.")
	}
}
