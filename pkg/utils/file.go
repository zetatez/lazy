package utils

import (
	"os"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

func GetFileBase(filePath string) string {
	return path.Base(filePath)
}

func GetFileExt(filePath string) string {
	return strings.ToLower(strings.TrimPrefix(path.Ext(filePath), "."))
}

func GetFilePrefix(filePath string) string {
	base := path.Base(filePath)
	ext := path.Ext(filePath)
	if ext == "" {
		return base
	}
	return base[:len(base)-len(ext)]
}

func GetFileParent(filePath string) string {
	return path.Dir(filePath)
}

func GetFileMimeType(filePath string) string {
	m, err := mimetype.DetectFile(filePath)
	if err != nil {
		return ""
	}
	ls := strings.SplitN(m.String(), ";", 2)
	if len(ls) > 0 {
		return ls[0]
	}
	return ""
}

func GetFileSize(filePath string) int64 {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	return info.Size()
}
