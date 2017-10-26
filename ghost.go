package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const (
	USER_FILEMODE = 0644
	GHOST_MESSAGE = "ghosted:"
)

var (
	BLANK = []byte{}
)

func ghost(pathname string, file os.FileInfo) {
	p := path.Join(pathname, file.Name())

	if file.IsDir() {
		fs, err := ioutil.ReadDir(p)
		if err != nil {
			fmt.Println("error on:", file.Name())
			return
		}
		for _, v := range fs {
			ghost(p, v)
		}
	} else {
		ioutil.WriteFile(p, BLANK, USER_FILEMODE)
		fmt.Println(GHOST_MESSAGE, p)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Need a filepath")
		return
	}

	root := os.Args[1]
	if fs, err := ioutil.ReadDir(root); err == nil {
		for _, v := range fs {
			ghost(root, v)
		}
	} else {
		fmt.Println(err)
	}
}
