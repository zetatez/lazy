package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/zetatez/lazy/config"
	"github.com/zetatez/lazy/utils/sugar"
)

const (
	VERSION = "0.0.1"
)

const (
	OptionView   = "view"
	OptionOpen   = "open"
	OptionExec   = "exec"
	OptionRename = "rename"
	OptionCopy   = "copy"
	OptionDelete = "delete"
)

var Options = map[string]bool{
	OptionView:   true,
	OptionOpen:   true,
	OptionExec:   true,
	OptionRename: true,
	OptionCopy:   true,
	OptionDelete: true,
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

func (l *Lazy) exec(cmd string) (err error) {
	fmt.Println("bash -c ", cmd)
	c := exec.Command("bash", "-c", cmd)
	c.Stdout, c.Stderr = os.Stdout, os.Stderr
	return c.Run()
}

func (l *Lazy) View() {
	cfg1, err1 := config.LoadConfig(
		path.Join("view", "ext", l.ext),
	)
	if err1 == nil {
		for _, cmd := range cfg1.Cmds {
			err := l.exec(
				fmt.Sprintf(`%s '%s'`, cmd, l.filePath),
			)
			if err == nil {
				return
			}
		}
	}

	cfg2, err2 := config.LoadConfig(
		path.Join("view", "mime", l.mimetype),
	)
	if err2 == nil {
		for _, cmd := range cfg2.Cmds {
			err := l.exec(
				fmt.Sprintf(`%s '%s'`, cmd, l.filePath),
			)
			if err == nil {
				return
			}
		}
	}
}

func (l *Lazy) Open() {
	cfg1, err1 := config.LoadConfig(
		path.Join("open", "ext", l.ext),
	)
	if err1 == nil {
		for _, cmd := range cfg1.Cmds {
			err := l.exec(
				fmt.Sprintf(`%s '%s'`, cmd, l.filePath),
			)
			if err == nil {
				return
			}
		}
	}

	cfg2, err2 := config.LoadConfig(
		path.Join("open", "mime", l.mimetype),
	)
	if err2 == nil {
		for _, cmd := range cfg2.Cmds {
			err := l.exec(
				fmt.Sprintf(`%s '%s'`, cmd, l.filePath),
			)
			if err == nil {
				return
			}
		}
	}
}

func (l *Lazy) Exec() {
	cfg1, err1 := config.LoadConfig(
		path.Join("exec", "ext", l.ext),
	)
	if err1 == nil {
		for _, cmd := range cfg1.Cmds {
			err := l.exec(
				fmt.Sprintf(`%s '%s'`, cmd, l.filePath),
			)
			if err == nil {
				return
			}
		}
	}
}

func (l *Lazy) Rename() {
	var newFileName string
	parent := sugar.GetFileParent(l.filePath)
	fmt.Printf("rename %s -> %s", l.filePath, parent)
	fmt.Scanf("%s", &newFileName)
	if newFileName == "" {
		newFileName = sugar.GetFileBase(l.filePath) + ".bk"
	}
	newFilePath := path.Join(parent, newFileName)
	os.Rename(l.filePath, newFilePath)
}

func (l *Lazy) Delete() {
	fmt.Printf("rm -rf %s\n", l.filePath)
	os.Remove(l.filePath)
}

func (l *Lazy) Copy() {
	parent := sugar.GetFileParent(l.filePath)

	fmt.Printf("copy %s -> %s", l.filePath, parent)

	var newFileName string
	fmt.Scanf("%s", &newFileName)
	if newFileName == "" {
		newFileName = sugar.GetFileBase(l.filePath) + ".bk"
	}
	newFilePath := path.Join(parent, newFileName)

	src, err := os.Open(l.filePath)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(newFilePath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return
	}
	defer dst.Close()

	io.Copy(dst, src)
}

func (l *Lazy) Version() {
	fmt.Println(VERSION)
}

func (l *Lazy) Help() {
	docs := `
	NAME
		lazy - A cli tool that greatly improves your work efficiency.

	SYNOPSIS
		lazy -v
		lazy -h
		lazy -o option -f file

	DESCRIPTION
		lazy is a tool for cli to open, exec, copy, rename, delete file automatically.

	OPTIONS
		-view     view file with your default setting
		-open     open file with your default setting
		-exec     exec script with your default setting
		-copy     copy file
		-rename   rename file
		-delete   delete file.

	BUGS
		Send all bug reports with a patch to zetatez@icloud.com.
	"`
	fmt.Println(docs)
}

func main() {
	h := flag.Bool("h", false, "help")
	v := flag.Bool("v", false, "version")
	option := flag.String("o", "", "option, available options: view, open, exec, copy, rename, delete")
	filePath := flag.String("f", "", "filePath")
	flag.Parse()

	if *h {
		NewLazy("").Help()
		return
	}
	if *v {
		NewLazy("").Version()
		return
	}
	if _, ok := Options[*option]; !ok {
		NewLazy("").Help()
		return
	}
	if *filePath == "" {
		return
	}

	if exists, _ := sugar.IsFileExists(*filePath); !exists {
		fmt.Println("file not exists")
		return
	}

	lazy := NewLazy(*filePath)
	switch *option {
	case OptionView:
		lazy.View()
	case OptionOpen:
		lazy.Open()
	case OptionExec:
		lazy.Exec()
	case OptionRename:
		lazy.Rename()
	case OptionCopy:
		lazy.Copy()
	case OptionDelete:
		lazy.Delete()
	default:
		NewLazy("").Help()
	}
}
