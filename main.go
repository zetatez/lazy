package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	yaml "gopkg.in/yaml.v2"
)

const (
	VERSION = "0.0.1"
)

var cfgFiles = []string{
	"lazy.yaml",
	path.Join(os.Getenv("HOME"), ".config", "lazy", "lazy.yaml"),
	path.Join(os.Getenv("HOME"), ".lazy.yaml"),
	"/etc/lazy.yaml",
}

type Yaml struct {
	View struct {
		Preset map[string]string `yaml:"preset"`
		Ext    map[string]string `yaml:"ext"`
		Mime   map[string]string `yaml:"mime"`
	} `yaml:"view"`
	Open struct {
		Preset map[string]string `yaml:"preset"`
		Ext    map[string]string `yaml:"ext"`
		Mime   map[string]string `yaml:"mime"`
	} `yaml:"open"`
	Exec struct {
		Preset map[string]string `yaml:"preset"`
		Ext    map[string]string `yaml:"ext"`
		Mime   map[string]string `yaml:"mime"`
	} `yaml:"exec"`
}

var cfg *Yaml

func loadCfg() (err error) {
	for _, f := range cfgFiles {
		isExists, err := isFileExists(f)
		if err != nil {
			return err
		}
		if !isExists {
			continue
		}

		fbyte, err := os.ReadFile(f)
		if err != nil {
			return err
		}
		if err = yaml.Unmarshal(fbyte, &cfg); err != nil {
			return err
		}
		break
	}

	if cfg == nil {
		return fmt.Errorf("no config file was found")
	}
	return nil
}

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
}

func NewLazy(filePath string) *Lazy {
	return &Lazy{filePath: filePath}
}

func (l *Lazy) base() (base string) {
	base = path.Base(l.filePath)
	return
}

func (l *Lazy) ext() (ext string) {
	ext = path.Ext(l.filePath)
	return
}

func (l *Lazy) extx() (ext string) {
	ext = l.ext()
	if ext != "" {
		ext = ext[1:]
	}
	return
}

func (l *Lazy) prefix() (prefix string) {
	base := l.base()
	ext := l.ext()
	prefix = base[:len(base)-len(ext)]
	return
}

func (l *Lazy) path() (path string) {
	path = l.filePath[:len(l.filePath)-len(l.base())]
	return
}

func (l *Lazy) mimeType() (mimeType string) {
	m, err := mimetype.DetectFile(l.filePath)
	if err != nil {
		return ""
	}
	ls := strings.Split(m.String(), ";")
	if len(ls) > 0 {
		mimeType = ls[0]
	}
	return mimeType
}

func (l *Lazy) View() {
	ext := l.extx()
	cmd := ""
	fakeCmd1, ok1 := cfg.View.Ext[ext]
	fakeCmd2, ok2 := cfg.View.Preset[fakeCmd1]
	if ok1 && !ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd1, l.filePath)
		run(cmd)
		return
	}
	if ok1 && ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd2, l.filePath)
		run(cmd)
		return
	}

	mimeType := l.mimeType()
	fakeCmd1, ok1 = cfg.View.Mime[mimeType]
	fakeCmd2, ok2 = cfg.View.Preset[fakeCmd1]
	if ok1 && !ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd1, l.filePath)
		run(cmd)
		return
	}
	if ok1 && ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd2, l.filePath)
		run(cmd)
		return
	}
}

func (l *Lazy) Open() {
	ext := l.extx()
	cmd := ""
	fakeCmd1, ok1 := cfg.Open.Ext[ext]
	fakeCmd2, ok2 := cfg.Open.Preset[fakeCmd1]
	if ok1 && !ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd1, l.filePath)
		run(cmd)
		return
	}
	if ok1 && ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd2, l.filePath)
		run(cmd)
		return
	}

	mimeType := l.mimeType()
	fakeCmd1, ok1 = cfg.Open.Mime[mimeType]
	fakeCmd2, ok2 = cfg.Open.Preset[fakeCmd1]
	if ok1 && !ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd1, l.filePath)
		run(cmd)
		return
	}
	if ok1 && ok2 {
		cmd = fmt.Sprintf(`%s '%s'`, fakeCmd2, l.filePath)
		run(cmd)
		return
	}
}

func (l *Lazy) Exec() {
	ext := l.extx()
	cmd := ""
	fakeCmd1, ok1 := cfg.Exec.Ext[ext]
	fakeCmd2, ok2 := cfg.Exec.Preset[fakeCmd1]
	if ok1 && !ok2 {
		if strings.Contains(fakeCmd1, "{}") {
			cmd = strings.ReplaceAll(fakeCmd1, "{}", fmt.Sprintf(`"%s"`, l.filePath))
		} else {
			cmd = fakeCmd1
		}
		run(cmd)
		return
	}
	if ok1 && ok2 {
		if strings.Contains(fakeCmd2, "{}") {
			cmd = strings.ReplaceAll(fakeCmd2, "{}", fmt.Sprintf(`"%s"`, l.filePath))
		} else {
			cmd = fakeCmd2
		}
		run(cmd)
		return
	}

	mimeType := l.mimeType()
	fakeCmd1, ok1 = cfg.Exec.Mime[mimeType]
	fakeCmd2, ok2 = cfg.Exec.Preset[fakeCmd1]
	if ok1 && !ok2 {
		if strings.Contains(fakeCmd1, "{}") {
			cmd = strings.ReplaceAll(fakeCmd1, "{}", fmt.Sprintf(`"%s"`, l.filePath))
		} else {
			cmd = fakeCmd1
		}
		run(cmd)
		return
	}
	if ok1 && ok2 {
		if strings.Contains(fakeCmd1, "{}") {
			cmd = strings.ReplaceAll(fakeCmd2, "{}", fmt.Sprintf(`"%s"`, l.filePath))
		} else {
			cmd = fakeCmd1
		}
		run(cmd)
		return
	}
}

func (l *Lazy) Rename() {
	var newFileName string
	fmt.Printf("rename %s -> %s", l.filePath, l.path())
	fmt.Scanf("%s", &newFileName)
	if newFileName == "" {
		newFileName = l.base() + ".bk"
	}
	newFilePath := path.Join(l.path(), newFileName)
	os.Rename(l.filePath, newFilePath)
}

func (l *Lazy) Delete() {
	fmt.Printf("rm -rf %s\n", l.filePath)
	os.Remove(l.filePath)
}

func (l *Lazy) Copy() {
	var newFileName string
	fmt.Printf("copy %s -> %s", l.filePath, l.path())
	fmt.Scanf("%s", &newFileName)
	if newFileName == "" {
		newFileName = l.base() + ".bk"
	}
	newFilePath := path.Join(l.path(), newFileName)

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

func isFileExists(filePath string) (exists bool, err error) {
	_, err = os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func run(cmd string) {
	fmt.Println("bash -c ", cmd)
	c := exec.Command("bash", "-c", cmd)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
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
	if *filePath == "" {
		NewLazy("").Help()
		return
	}
	if exists, _ := isFileExists(*filePath); !exists {
		fmt.Println("file not exists")
		return
	}

	if _, ok := Options[*option]; !ok {
		NewLazy("").Help()
		return
	}

	err := loadCfg()
	if err != nil {
		fmt.Println(err)
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
