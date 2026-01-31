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

const versionText = "0.1.0"

const startupText = `- file path: %s
  - ext: %s
  - mimetype: %s
- option: %s
`

const helpText = `NAME
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
	Report bugs to zetatez@icloud.com.`

type Action func(*Lazy, *Config) bool

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

func (l *Lazy) VIEW(cfg *Config) bool {
	return l.runCmd(cfg, "view")
}

func (l *Lazy) OPEN(cfg *Config) bool {
	return l.runCmd(cfg, "open")
}

func (l *Lazy) EXEC(cfg *Config) bool {
	return l.runCmd(cfg, "exec")
}

func (l *Lazy) runCmd(cfg *Config, action string) bool {
	var cmds []string

	cmds = append(cmds, cfg.getCmds(action, "ext", l.ext)...)
	cmds = append(cmds, cfg.getCmds(action, "mimetype", l.mimetype)...)
	cmds = append(cmds, cfg.getCmds(action, "ext", "default")...)
	cmds = append(cmds, cfg.getCmds(action, "mimetype", "default")...)

	for _, cmd := range cmds {
		finalCmd := fmt.Sprintf(`%s '%s'`, cmd, escapeShellArg(l.filePath))
		if err := l.exec(finalCmd); err != nil {
			fmt.Fprintf(os.Stderr, "%s failed: %v\n", cmd, err)
			continue
		}
		fmt.Println(finalCmd)
		return true
	}
	return false
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
	fmt.Println(helpText)
}

func (l *Lazy) PrintVersion() {
	fmt.Println(versionText)
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
	action, ok := Options[*option]
	if !ok {
		fmt.Println("Error: invalid option. Use -h for help.")
	}
	if !action(lazy, cfg) {
		fmt.Printf(startupText, *filePath, lazy.ext, lazy.mimetype, *option)
	}
}
