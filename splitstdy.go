package main

import (
	"strings"
	"fmt"
)

func main() {
	uri := "/?cycles=12"
	cycle := strings.Split(uri, "/?cycles=")[1]
	fmt.Println(cycle)
}
