package src

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func HashObject(root string, file string) {
	if !isFile(file) {
		fmt.Println("Error:" + file + "is not file!")
		return
	}
	content := getContent(file)
	header := fmt.Sprintf("%s %d\u0000", "blob", len(content))
	data := append([]byte(header), content...)

	pwd := sha1.Sum(data)
	pwdstr := fmt.Sprintf("%x", pwd)
	fmt.Println(pwdstr)

	prefix := pwdstr[:2]
	postfix := pwdstr[2:]

}

func getContent(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return content
}

func findRoot(path string) string {
	for {
		path = filepath.Join(path, "..")
		fileName, err := filepath.Glob(filepath.Join(path, "*"))
		if err != nil {
			fmt.Println("Error:", err)
			return ""
		}
		for i := range fileName {
			if fileName[i] == ".tdf" {
				res := filepath.Join(path, ".tdf")
				return res
			}
		}
	}
}

func isFile(f string) bool {
	stat, err := os.Stat(f)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}
