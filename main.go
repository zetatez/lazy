package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/zetatez/lazy/config"
	"github.com/zetatez/lazy/sugar"
)

const VERSION = "0.0.1"

var Options = map[string]func(*Lazy){
	"view": (*Lazy).View,
	"open": (*Lazy).Open,
	"exec": (*Lazy).Exec,
	"mv":   (*Lazy).Mv,
	"cp":   (*Lazy).Cp,
	"rm":   (*Lazy).Rm,
}

type Lazy struct {
	filePath string
	mimetype string
	ext      string
}

func NewLazy(filePath string) *Lazy {
	return &Lazy{
		filePath: filePath,
		mimetype: sugar.GetFileMimeType(filePath),
		ext:      sugar.GetFileExt(filePath),
	}
}

func (l *Lazy) exec(cmd string) error {
	fmt.Println("Executing:", cmd)
	c := exec.Command("bash", "-c", cmd)
	c.Stdout, c.Stderr = os.Stdout, os.Stderr
	return c.Run()
}

func (l *Lazy) runCmd(configPath string) (err error) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return err
	}
	for _, cmd := range cfg.Cmds {
		if err = l.exec(fmt.Sprintf(`%s '%s'`, cmd, l.filePath)); err == nil {
			return nil
		} else {
			sugar.Notify(err)
		}
	}
	return err
}

func (l *Lazy) View() {
	if l.runCmd(path.Join("view", "ext", l.ext)) == nil {
		return
	}
	l.runCmd(path.Join("view", "mime", l.mimetype))
}

func (l *Lazy) Open() {
	if l.runCmd(path.Join("open", "ext", l.ext)) == nil {
		return
	}
	l.runCmd(path.Join("open", "mime", l.mimetype))
}

func (l *Lazy) Exec() {
	l.runCmd(path.Join("exec", "ext", l.ext))
}

func (l *Lazy) Mv() {
	parent := sugar.GetFileParent(l.filePath)
	fmt.Printf("mv %s %s", l.filePath, parent)
	var newFileName string
	fmt.Scanf("%s", &newFileName)
	if newFileName == "" {
		newFileName = sugar.GetFileBase(l.filePath) + ".bk"
	}
	newFilePath := path.Join(parent, newFileName)
	if err := os.Rename(l.filePath, newFilePath); err != nil {
		fmt.Println("Error moving file:", err)
		return
	}
}

func (l *Lazy) Rm() {
	fmt.Printf("rm %s\n", l.filePath)
	if err := os.Remove(l.filePath); err != nil {
		fmt.Println("Error removing file:", err)
		return
	}
}

func (l *Lazy) Cp() {
	parent := sugar.GetFileParent(l.filePath)
	fmt.Printf("cp %s %s", l.filePath, parent)
	var newFileName string
	fmt.Scanf("%s", &newFileName)
	if newFileName == "" {
		newFileName = sugar.GetFileBase(l.filePath) + ".bk"
	}
	newFilePath := path.Join(parent, newFileName)

	src, err := os.Open(l.filePath)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(newFilePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println("Error copying file:", err)
		return
	}
}

func (l *Lazy) Version() {
	fmt.Println(VERSION)
}

func (l *Lazy) Help() {
	fmt.Println(`
NAME
	lazy - A CLI tool that improves your work efficiency.

SYNOPSIS
	lazy -v
	lazy -h
	lazy -o <OPTION> -f <file>

DESCRIPTION
	lazy is a tool for CLI to view, open, execute, copy, move, or remove files automatically.

OPTIONS
	-view    View file with your default setting.
	-open    Open file with your default setting.
	-exec    Execute script with your default setting.
	-cp      Copy file.
	-mv      Move file.
	-rm      Remove file.

BUGS
	Report bugs to zetatez@icloud.com.
	`)
}

func main() {
	h := flag.Bool("h", false, "help")
	v := flag.Bool("v", false, "version")
	option := flag.String("o", "", "operation (view, open, exec, cp, mv, rm)")
	filePath := flag.String("f", "", "file path")
	flag.Parse()

	if *h {
		NewLazy("").Help()
		return
	}
	if *v {
		NewLazy("").Version()
		return
	}

	if *filePath == "" || !sugar.IsFileExists(*filePath) {
		fmt.Println("Error: File does not exist.")
		return
	}

	lazy := NewLazy(*filePath)
	if action, ok := Options[*option]; ok {
		action(lazy)
	} else {
		fmt.Println("Invalid option. Use -h for help.")
	}
}
