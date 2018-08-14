package main

import (
	"net/http"
	"io"
	"os"
	"fmt"
	"strings"
)

func main() {
	resp, err := http.Head("http://baidu.com/")
	resp1, err := http.Post("http://baidu.com/", "text", strings.NewReader("name=cjb"))
	fmt.Println(resp)
	//contentTypes, ok := r.Header["Content-Type"]
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	defer resp1.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
	io.Copy(os.Stdout, resp1.Body)
	fmt.Println(resp1.Header)
}
