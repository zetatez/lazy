package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"

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

func (l *Lazy) exec(cmd string) {
	fmt.Println("bash -c ", cmd)
	c := exec.Command("bash", "-c", cmd)
	c.Stdout, c.Stderr = os.Stdout, os.Stderr
	c.Run()
}

// func (l *Lazy) View() {
// 	cmd := ""
// 	fakeCmd1, ok1 := config.GetConfig().View.Ext[sugar.GetFileExt(l.filePath)]
// 	fakeCmd2, ok2 := config.GetConfig().View.Preset[fakeCmd1]
// 	if ok1 && !ok2 {
// 		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd1, l.filePath)
// 		l.exec(cmd)
// 		return
// 	}
// 	if ok1 && ok2 {
// 		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd2, l.filePath)
// 		l.exec(cmd)
// 		return
// 	}
// 	fakeCmd1, ok1 = config.GetConfig().View.Mime[sugar.GetFileMimeType(l.filePath)]
// 	fakeCmd2, ok2 = config.GetConfig().View.Preset[fakeCmd1]
// 	if ok1 && !ok2 {
// 		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd1, l.filePath)
// 		l.exec(cmd)
// 		return
// 	}
// 	if ok1 && ok2 {
// 		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd2, l.filePath)
// 		l.exec(cmd)
// 		return
// 	}
// }

func (l *Lazy) getCfg(filePath string) (cfg string) {
	return cfg
}

func (l *Lazy) View() {
	cmd := ""
	fakeCmd1, ok1 := config.GetConfig().View.Ext[sugar.GetFileExt(l.filePath)]
	fakeCmd2, ok2 := config.GetConfig().View.Preset[fakeCmd1]
	if ok1 && !ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd1, l.filePath)
		l.exec(cmd)
		return
	}
	if ok1 && ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd2, l.filePath)
		l.exec(cmd)
		return
	}
	fakeCmd1, ok1 = config.GetConfig().View.Mime[sugar.GetFileMimeType(l.filePath)]
	fakeCmd2, ok2 = config.GetConfig().View.Preset[fakeCmd1]
	if ok1 && !ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd1, l.filePath)
		l.exec(cmd)
		return
	}
	if ok1 && ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd2, l.filePath)
		l.exec(cmd)
		return
	}
}

func (l *Lazy) Open() {
	cmd := ""
	fakeCmd1, ok1 := config.GetConfig().Open.Ext[sugar.GetFileExt(l.filePath)]
	fakeCmd2, ok2 := config.GetConfig().Open.Preset[fakeCmd1]
	if ok1 && !ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd1, l.filePath)
		l.exec(cmd)
		return
	}
	if ok1 && ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd2, l.filePath)
		l.exec(cmd)
		return
	}
	fakeCmd1, ok1 = config.GetConfig().Open.Mime[sugar.GetFileMimeType(l.filePath)]
	fakeCmd2, ok2 = config.GetConfig().Open.Preset[fakeCmd1]
	if ok1 && !ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd1, l.filePath)
		l.exec(cmd)
		return
	}
	if ok1 && ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd2, l.filePath)
		l.exec(cmd)
		return
	}
}

func (l *Lazy) Exec() {
	cmd := ""
	fakeCmd1, ok1 := config.GetConfig().Exec.Ext[sugar.GetFileExt(l.filePath)]
	fakeCmd2, ok2 := config.GetConfig().Exec.Preset[fakeCmd1]
	if ok1 && !ok2 {
		if strings.Contains(fakeCmd1, "{}") {
			cmd = strings.ReplaceAll(fakeCmd1, "{}", fmt.Sprintf(`"%s"`, l.filePath))
		} else {
			cmd = fakeCmd1
		}
		l.exec(cmd)
		return
	}
	if ok1 && ok2 {
		if strings.Contains(fakeCmd2, "{}") {
			cmd = strings.ReplaceAll(fakeCmd2, "{}", fmt.Sprintf(`"%s"`, l.filePath))
		} else {
			cmd = fakeCmd2
		}
		l.exec(cmd)
		return
	}
	fakeCmd1, ok1 = config.GetConfig().Exec.Mime[sugar.GetFileMimeType(l.filePath)]
	fakeCmd2, ok2 = config.GetConfig().Exec.Preset[fakeCmd1]
	if ok1 && !ok2 {
		if strings.Contains(fakeCmd1, "{}") {
			cmd = strings.ReplaceAll(fakeCmd1, "{}", fmt.Sprintf(`"%s"`, l.filePath))
		} else {
			cmd = fakeCmd1
		}
		l.exec(cmd)
		return
	}
	if ok1 && ok2 {
		if strings.Contains(fakeCmd1, "{}") {
			cmd = strings.ReplaceAll(fakeCmd2, "{}", fmt.Sprintf(`"%s"`, l.filePath))
		} else {
			cmd = fakeCmd1
		}
		l.exec(cmd)
		return
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
	var newFileName string
	parent := sugar.GetFileParent(l.filePath)
	fmt.Printf("copy %s -> %s", l.filePath, parent)
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

	// err := config.GetConfig().LoadCfg()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

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
