package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")
	for i := 1; ;i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("hello")
	}
	time.Sleep(time.Second * 2)

	fmt.Println("World")
}
