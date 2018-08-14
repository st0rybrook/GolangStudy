package main

import (
	"runtime"
	"bytes"
	"strconv"
	"sync/atomic"
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()
		//t := GetGID()
		//fmt.Println(t)
		i := int32(1)
		b := atomic.AddInt32(&i, 1)
		fmt.Println(i)
		i++
		fmt.Println(b)
	}()

	wg.Wait()
	printStack()
}
func GetGID() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	gid, _ = strconv.ParseUint(string(b), 10, 64)
	return
}
func printStack() {
	//buf := make([]byte, 1<<19)
	//runtime.Stack(buf, true)
	//fmt.Printf("\n%s", buf)
	//if err := recover(); err != nil {
		stack := make([]byte, 1024*8)
		stack = stack[:runtime.Stack(stack, false)]
		//f := "[PANIC] %s\n%s"
		//VPNGatewayCommonLibs.ERRORF(f, err, stack)
		fmt.Printf("\n%s", stack)
	//}

}
