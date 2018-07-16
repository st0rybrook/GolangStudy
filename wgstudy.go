package main

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			ch<-i
			wg.Done()
		}(i)
	}

	for i :=0; i<5; i++ {
		v:=<-ch
		log.Println(v)
	}
	wg.Wait()
	log.Println("exit")


}