package main

import "fmt"

func main() {
	s := "tangshenzheng"
	b := []byte(s)
	b[0] = 103
	fmt.Println(b)
	s = string(b)
	fmt.Println(s)
}
