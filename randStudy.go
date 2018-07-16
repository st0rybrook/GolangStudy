package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	eipList := []string{"eip1", "eip2", "eip3", "eip4"}
	id := selectBindEipId(eipList)
	fmt.Println(id)
	fmt.Println(eipList)
}

//eiplist至少两个eip-id
func selectBindEipId(eipList []string) string {
	switch len(eipList) {
	case 0:
		{
			panic("error")
		}
	case 1:
		{
			panic("error")
		}
	case 2:
		{
			return eipList[1]
		}
	default:
		{
			rand.Seed(time.Now().Unix())
			rnd := rand.Intn(len(eipList) - 1)
			eipList[0], eipList[rnd] = eipList[rnd], eipList[0]
			return eipList[0]
		}
	}

}
