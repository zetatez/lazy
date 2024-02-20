package sugar

import (
	"os"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func IsFileExists(filePath string) (exists bool, err error) {
	if _, err = os.Stat(filePath); err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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
