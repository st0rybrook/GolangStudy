package main

import (
	"fmt"
	"os/exec"

	"uframework/log"
)

func main() {
	cmdline := fmt.Sprintf("lspci")
	output, err := runCmd(cmdline, "5")
	if err != nil {
		uflog.ERRORF("get [namespace:%s] tunnel status error!")
	}
	fmt.Println(output)
}
func runCmd(command, timeout string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("timeout %s %s", timeout, command))
	out, err := cmd.CombinedOutput()
	uflog.DEBUGF("%q", string(out))
	if err != nil {
		uflog.ERRORF("%s", err.Error())
	}
	return string(out), err
}
