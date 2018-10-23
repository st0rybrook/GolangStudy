package main

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"gitlab.ucloudadmin.com/unetwork/vpngw_common/libs"
)

func main() {
	errMsg:=""
	errMsg = fmt.Sprintf("%s Failed to del route rule: %s -> %s via %s",errMsg,"123","456","678")
	errMsg=fmt.Sprintf("%s \n aaa",errMsg)
	fmt.Println(errors.New(errMsg))
	log.Debugf(errMsg)
	VPNGatewayCommonLibs.LOG(nil,"DeleteVPNGateway").Debug("VPN Gateway is in use, failed to delete")

}
