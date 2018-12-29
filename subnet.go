package main

import (
	"fmt"
	"net"
)

func main() {
	local := []string{"192.168.0.0/16"}
	remote := []string{"192.168.1.0/24"}
	err := checkIPSecSubnets(local, remote)
	fmt.Println(err.Error())

}
func checkIPSecSubnets(localSubnets []string, remoteSubnets []string) error {
	// parse string to subnet
	var localNets, remoteNets []*net.IPNet
	for _, str := range localSubnets {
		_, n, err := net.ParseCIDR(str)
		if err != nil {
			return err
		}
		fmt.Println(n)
		localNets = append(localNets, n)
	}
	for _, str := range remoteSubnets {
		_, n, err := net.ParseCIDR(str)
		if err != nil {
			return err
		}
		remoteNets = append(remoteNets, n)
	}

	contains := func(net1 *net.IPNet, net2 *net.IPNet) bool {
		if net1.Contains(net2.IP) {
			return true
		}
		if net2.Contains(net1.IP) {
			return true
		}
		return false
	}
	// check localsubnets
	for i := 0; i < len(localNets)-1; i++ {
		for j := i + 1; j < len(localNets); j++ {
			if contains(localNets[i], localNets[j]) {
				return fmt.Errorf("localNets: %s, %s not valid", localNets[i].String(), localNets[j].String())
			}
		}
	}
	// check remotesubnets
	for i := 0; i < len(remoteNets)-1; i++ {
		for j := i + 1; j < len(remoteNets); j++ {
			if contains(remoteNets[i], remoteNets[j]) {
				return fmt.Errorf("remoteNets: %s, %s not valid", remoteNets[i].String(), remoteNets[j].String())
			}
		}
	}
	// check local and remote
	for _, net1 := range localNets {
		for _, net2 := range remoteNets {
			if contains(net1, net2) {
				return fmt.Errorf("localNets: %s, remoteNets: %s not valid", net1.String(), net2.String())
			}
		}
	}
	return nil
}
