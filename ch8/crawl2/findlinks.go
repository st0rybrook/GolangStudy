// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
//带跳转次数的url树
type urlTree struct {
	urllist []string
	depth   int
}

func crawl(url string, depth int) urlTree {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return urlTree{list, depth + 1}
}

//!-sema

//!+

func main() {
	worklist := make(chan urlTree)
	var n int // number of pending sends to worklist
	var depth int
	var depthLimit int = 1
	// Start with the command-line arguments.
	n++
	go func() { worklist <- urlTree{os.Args[1:], depth} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if list.depth > depthLimit {
			continue
		}
		for _, link := range list.urllist {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string, depth int) {
					worklist <- crawl(link, list.depth)
				}(link, list.depth)
			}
		}
	}
}

//!-
