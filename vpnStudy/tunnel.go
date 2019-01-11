package main

import (
	"fmt"
	"strings"
)

func main() {
	output := `Security Associations (1 up, 0 connecting):
	vpntunnel-ugtmlf_1[1]: ESTABLISHED 16 hours ago, 120.132.29.114[120.132.29.114]...117.50.61.96[117.50.61.96]
	vpntunnel-ugtmlf_1{22}:  INSTALLED, TUNNEL, reqid 1, ESP SPIs: cfa04e50_i ccaf1ffa_o
	vpntunnel-ugtmlf_1{22}:   192.168.0.0/30 === 172.24.0.0/24`
	inbytes:=`src 120.132.29.114 dst 117.50.61.96
	proto esp spi 0xce15ff2d(3457548077) reqid 1(0x00000001) mode tunnel
	replay-window 0 seq 0x00000000 flag af-unspec (0x00100000)
	auth-trunc hmac(sha1) 0x3a13962d519e77dfdf745c74a83b9c7588123fe1 (160 bits) 96
	enc cbc(aes) 0x6abc7e8ea43d365d42a9b8d71c7db9be (128 bits)
	anti-replay context: seq 0x0, oseq 0x22b, bitmap 0x00000000
	lifetime config:
	  limit: soft (INF)(bytes), hard (INF)(bytes)
	  limit: soft (INF)(packets), hard (INF)(packets)
	  expire add: soft 2612(sec), hard 3600(sec)
	  expire use: soft 0(sec), hard 0(sec)
	lifetime current:
	  46620(bytes), 555(packets)
	  add 2019-01-10 11:11:54 use 2019-01-10 11:11:54
	stats:
	  replay-window 0 replay 0 failed 0`
	  line1:=strings.Split(inbytes, "\n")[0]
	  fmt.Println(strings.Split(line1," ")[1])
	  fmt.Println(line1)
	//取Security以后的字符串,去除最后一个换行,逐行扫描
	index := strings.Index(output, "Security")
	out := substr(output, index, len(output))
	out = strings.Trim(out, "\n")
	lines := strings.Split(out, "\n")
	// ignore tunnel with no SA established
	str := strings.Replace(lines[1], " ", "", -1)
	fmt.Println(str)
	remoteIP:=strings.Split(strings.Split(str,"...")[1],"[")[0]
	fmt.Println(remoteIP)
	fmt.Println(lines)
	for _, line := range lines {
		if strings.Contains(line, "SPI") {
			line = strings.Trim(line, " ")
			tunnelName := strings.Split(line, "_")[0]
			spiId := strings.Split(line, ":")[2]
			spiId = strings.Trim(spiId, " ")
			spiInId := strings.Split(spiId, " ")[0]
			spiOutId := strings.Split(spiId, " ")[1]

			fmt.Println(tunnelName)
			fmt.Println(spiInId)
			fmt.Println(spiOutId)
		}

	}
}
func substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	end = start + length
	if length < 0 {
		end = rl
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
