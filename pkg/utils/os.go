package utils

import (
	"os"
	"runtime"
)

func GetOSType() string {
	return runtime.GOOS
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
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

func IsDirExists(path string) bool {
	if Exists(path) && IsDir(path) {
		return true
	}
	return false
}

func IsFileExists(path string) bool {
	if Exists(path) && IsFile(path) {
		return true
	}
	return false
}

func IsSymlink(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeSymlink != 0
}
