package src

import (
	"crypto/sha1"
	"os"
	"path/filepath"
)

func HashObject(file string) {
	if isFile(file) {

	}
}

func getContent(path string) {

}

func isFile(f string) bool {
	stat, err := os.Stat(f)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}
