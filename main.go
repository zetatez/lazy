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

const (
	VERSION = "0.0.1"
)

const (
	OptionVIEW = "view"
	OptionOPEN = "open"
	OptionEXEC = "exec"
	OptionMV   = "mv"
	OptionCP   = "cp"
	OptionRM   = "rm"
)

var Options = map[string]bool{
	OptionVIEW: true,
	OptionOPEN: true,
	OptionEXEC: true,
	OptionMV:   true,
	OptionCP:   true,
	OptionRM:   true,
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

func (l *Lazy) VIEW() {
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
			} else {
				sugar.Notify(err)
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
			} else {
				sugar.Notify(err)
			}
		}
	}
}

func (l *Lazy) OPEN() {
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
			} else {
				sugar.Notify(err)
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
			} else {
				sugar.Notify(err)
			}
		}
	}
}

func (l *Lazy) EXEC() {
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
			} else {
				sugar.Notify(err)
			}
		}
	}
}

func (l *Lazy) MV() {
	var newFileName string
	parent := sugar.GetFileParent(l.filePath)
	fmt.Printf("mv %s -> %s", l.filePath, parent)
	fmt.Scanf("%s", &newFileName)
	if newFileName == "" {
		newFileName = sugar.GetFileBase(l.filePath) + ".bk"
	}
	newFilePath := path.Join(parent, newFileName)
	os.Rename(l.filePath, newFilePath)
}

func (l *Lazy) RM() {
	fmt.Printf("rm -rf %s\n", l.filePath)
	os.Remove(l.filePath)
}

func (l *Lazy) CP() {
	parent := sugar.GetFileParent(l.filePath)

	fmt.Printf("cp %s -> %s", l.filePath, parent)

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
		lazy -o <OPTION> -f <file>

	DESCRIPTION
		lazy is a tool for cli to view, open, exec, cp, mv, rm file automatically.

	OPTION
		-view    view file   with your default setting
		-open    open file   with your default setting
		-exec    exec script with your default setting
		-cp      cp file
		-mv      mv file
		-rm      rm file

	BUGS
		Send bug report with a patch to zetatez@icloud.com.
	"`
	fmt.Println(docs)
}

func main() {
	h := flag.Bool("h", false, "help")
	v := flag.Bool("v", false, "version")
	option := flag.String("o", "", "option, available options: view, open, exec, cp, mv, rm")
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

	if !sugar.IsFileExists(*filePath) {
		fmt.Println("file not exists")
		return
	}

	lazy := NewLazy(*filePath)
	switch *option {
	case OptionVIEW:
		lazy.VIEW()
	case OptionOPEN:
		lazy.OPEN()
	case OptionEXEC:
		lazy.EXEC()
	case OptionMV:
		lazy.MV()
	case OptionCP:
		lazy.CP()
	case OptionRM:
		lazy.RM()
	default:
		NewLazy("").Help()
	}
}
