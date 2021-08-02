package main

import (
	"github.com/Wh0rigin/Then-diff/src"
	"os"
)

func main() {
	command := os.Args[1]
	switch command {
	case "init":
		path := os.Args[2]
		src.Init(path)
	case "cat-file":
	case "hash-object":
	}

}
