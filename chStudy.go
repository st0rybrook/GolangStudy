package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	//c <- 10
	fmt.Println("aaaa")
	//关闭之后可以读取
	close(c)
	//<-c
	v, ok := <-c // v=10,ok=true，虽然c关闭了，但是有数据，ok依然是true
	fmt.Println(v)
	fmt.Println(ok)
	//v, ok= <- c // v=0,ok=false，读失败了。
	//fmt.Println(v)
	//fmt.Println(ok)
}
