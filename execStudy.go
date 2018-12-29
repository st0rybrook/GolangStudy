package main
import (
	"fmt"
	"os/exec"
)

func checkNetInterfaceIsX710() bool{
	cmd := exec.Command("/bin/sh", "-c", "lspci | grep -i ethernet | grep -i X710")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return false
	}
	if string(out)=="" {
		return false
	}
	return true
}
