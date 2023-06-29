package sugar

import (
	"os"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func IsFileExists(filePath string) (exists bool, err error) {
	_, err = os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileBase(filePath string) (base string) {
	base = path.Base(filePath)
	return
}

func GetFileExt(filePath string) (ext string) {
	ext = path.Ext(filePath)
	return
}

func GetFileExtx(filePath string) (ext string) {
	ext = path.Ext(filePath)
	if ext == filePath {
		return filePath
	}
	if ext != "" && ext[0] == '.' {
		ext = ext[1:]
	}
	return
}

func GetFilePrefix(filePath string) (prefix string) {
	base := path.Base(filePath)
	ext := path.Ext(filePath)
	prefix = base[:len(base)-len(ext)]
	return
}

func GetFileParent(filePath string) (parent string) {
	parent = filePath[:len(filePath)-len(path.Base(filePath))]
	return
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
