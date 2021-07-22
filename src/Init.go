package src

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[tdf]")
}

func Init(path string) {
	//是否指定路径
	if path == "" {
		cur, _ := os.Getwd()
		mkdirfile(cur)
		return
	}
	//判断当前目录是否存在
	if ok, _ := pathExists(path); ok {
		mkdirfile(path)
	} else {
		fmt.Println("没有" + path + "文件目录!请核实文件目录后再次尝试")
	}

}

func mkdirfile(path string) {

}

func pathExists(path string) (bool, error) {
	stat, err := os.Stat(path)

	if err == nil {
		if !stat.IsDir() {
			log.Println(path + " is a file!")
			return false, nil
		}
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
