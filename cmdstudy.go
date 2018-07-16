package main

import (
	"fmt"
	"os/exec"
)

func main() {
	arg := "curl -I \"www.baidu.com\""

	cmd := exec.Command("sh", "-c",arg)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	cmd.Wait()
}
