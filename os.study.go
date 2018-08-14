package main

import (

	"io"
	"log"
	"os"
	"bufio"
)


func main() {

	//reader := bufio.NewReader(os.Stdin)
	//
	//result, err := reader.ReadString('\n')
	//if err != nil {
	//
	//	fmt.Println("read error:", err)
	//}
	//
	//
	//fmt.Println("result:", result)
	inputReader := bufio.NewReader(os.Stdin)
	//mustCopy(os.Stdout, inputReader)
	inputReader.WriteTo((os.Stdout))
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
