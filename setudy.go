package main

import (
	"fmt"
	"unsafe"
)

// 定义非空结构体
type S struct {
	a uint16
	b uint32
}

func main() {
	var s S
	fmt.Println(unsafe.Sizeof(s)) // prints 8, not 6
	var s2 struct{}
	fmt.Println(unsafe.Sizeof(s2)) // prints 0
	a := struct{}{}
	b := struct{}{}
	fmt.Println(a == b) // true
	fmt.Printf("%p, %p\n", &a, &b) // 0x55a988, 0x55a988

}