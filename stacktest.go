package main

import (
	"github.com/data_struct"
	"fmt"
)

func main() {
	var s data_struct.Stack
	s.Push(1)
	fmt.Println(s.Top())
	s.Push(2)
	fmt.Println(s.Top())
	eipList := []string{
		"eip-aeeiew",
		"eip-dekelg",
	}
	eipList[0], eipList[1] = eipList[1], eipList[0]
	fmt.Println(eipList)
}
