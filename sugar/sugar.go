package sugar

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func IsDirExists(path string) (exist bool) {
	if Exists(path) && IsDir(path) {
		return true
	}
	return false
}

func IsFileExists(path string) (exist bool) {
	if Exists(path) && IsFile(path) {
		return true
	}
	return false
}

func Exists(path string) (exist bool) {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func IsDir(path string) (isDir bool) {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func IsFile(path string) (isFile bool) {
	return !IsDir(path)
}

func GetFileBase(filePath string) (base string) {
	return path.Base(filePath)
}

func GetFileExt(filePath string) (ext string) {
	return strings.TrimPrefix(path.Ext(filePath), ".")
}

func GetFilePrefix(filePath string) (prefix string) {
	base := path.Base(filePath)
	ext := path.Ext(filePath)
	prefix = base[:len(base)-len(ext)]
	return prefix
}

func GetFileParent(filePath string) (parent string) {
	parent = filePath[:len(filePath)-len(path.Base(filePath))]
	return parent
}

func GetFileMimeType(filePath string) (mimeType string) {
	m, err := mimetype.DetectFile(filePath)
	if err != nil {
		return ""
	}
	ls := strings.Split(m.String(), ";")
	if len(ls) > 0 {
		mimeType = ls[0]
	}
	return mimeType
}

func Notify(msg ...interface{}) {
	NewExecService().RunScriptShell(fmt.Sprintf("notify-send '%v'", msg))
}
