package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	rd, err := ioutil.ReadDir("D:\\下载软件")
	for _, fi := range rd {
		fmt.Println("")
		fmt.Println(fi.Name())
		fmt.Println(fi.IsDir())
		fmt.Println(fi.Size())
		fmt.Println(fi.ModTime())
		fmt.Println(fi.Mode())
	}
	fmt.Println("")
	fmt.Println(err)
}
