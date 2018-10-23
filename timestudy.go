package main

import (
	"fmt"
	"time"
)

func main() {
	t:=time.Now()
	str_t := t.Format("2006-01-02 15:04:05" )
	fmt.Println(str_t)
	fmt.Println("hello")
	for i := 1;i<5 ;i++ {
		time.Sleep(time.Second * 1)
		fmt.Println("hello")
	}
	time.Sleep(time.Second * 2)

	fmt.Println("World")
	fmt.Println(t.Unix())
}
