package main

import (
	"fmt"
	"crypto/rand"
	"strings"
)
const (
	MAC_VPNGW_LAN = "52:90:0"
	MAC_VPNGW_WAN = "52:90:1"
)

func main() {
	buf := make([]byte, 4)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println("allocateMac Fail")
		return
	}
	fmt.Println(buf)
	mac := strings.ToUpper(fmt.Sprintf("%s%01x:%02x:%02x:%02x", MAC_VPNGW_LAN, buf[0]%16, buf[1], buf[2], buf[3]))
	macsuffix := strings.Join(strings.Split(mac, ":")[4:6], ":")
	fmt.Println(mac)
	fmt.Println(macsuffix)
}
