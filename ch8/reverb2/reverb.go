// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	"sync"
	"bufio"
)

//func echo(c net.Conn, shout string, delay time.Duration) {
//	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
//	time.Sleep(delay)
//	fmt.Fprintln(c, "\t", shout)
//	time.Sleep(delay)
//	fmt.Fprintln(c, "\t", strings.ToLower(shout))
//}

//!+
func handleConn(c net.Conn) {
	message := make(chan string)
	//tick := time.NewTicker(10 * time.Second)
	input := bufio.NewScanner(c)
	//input.Scan()
	var wg sync.WaitGroup
	//flag := input.Scan()
	//var flag bool = true
	wg.Add(1)
	/*
	1.启动一个goroutine,for死循环让他不能断掉
	select语句case判断两个channel
	一个是10秒后断掉连接
	另一个是接收标准输入后发送过来的channel，接收到值后，启动goroutinue输出

	2.for循环接收标准输入，接收到后发送给message的channel
	*/
	go func() {
		wg.Done()
		for {
			select {
			case <-time.After(time.Second * 10):
				c.Close()
			case mes := <-message:
				wg.Add(1)
				go func(c net.Conn, shout string, delay time.Duration) {
					defer wg.Done()
					fmt.Fprintln(c, "\t", strings.ToUpper(shout))
					time.Sleep(delay)
					fmt.Fprintln(c, "\t", shout)
					time.Sleep(delay)
					fmt.Fprintln(c, "\t", strings.ToLower(shout))
				}(c, mes, 1*time.Second)
			}
		}

	}()

	for input.Scan() {
		text := input.Text()
		message <- text
	}
	// NOTE: ignoring potential errors from input.Err()
	wg.Wait()
	//cw := c.(*net.TCPConn)
	//cw.CloseWrite()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
