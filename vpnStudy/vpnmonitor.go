package main

import (
	"flag"
	"fmt"
	"strconv"

	"uframework/common"

	"gitlab.ucloudadmin.com/unetwork/vpngw_common/libs"
	"gitlab.ucloudadmin.com/unetwork/vpngw_monitor/worker"
)

var (
	confFile    = flag.String("c", "", "configuration file,json format")
	appName     = flag.String("a", "", "application name")
	confService = flag.String("s", "", "config service address,http server address")
	regionId    = flag.String("r", "", "region id")
	AppConfig    VPNGatewayCommonLibs.Config
	RegionConfig VPNGatewayCommonLibs.Config
)


func main() {
	// 解释命令行选项
	//ufcommon.ProcessOptions()
	//ufcommon.DumpOptions()
	VPNGatewayCommonLibs.ProcessOptions()
	VPNGatewayCommonLibs.DumpOptions()
	// 处理配置
	option, err := VPNGatewayCommonLibs.GetOption("c")
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg := new(VPNGatewayCommonLibs.Config)
	if err = cfg.LoadConfigFromFile(option); err != nil {
		fmt.Println("Load Config File fail,", err)
		return
	}
	cfg.DumpConfigContent()
	regionId, err := VPNGatewayCommonLibs.GetOption("r")
	if err != nil {
		fmt.Println(err)
		return
	}
	myRegionId, err := strconv.Atoi(regionId)
	if err != nil {
		fmt.Println(err)
		return
	}
	if myRegionId == 0 {
		fmt.Println("RegionId can not be zero ...")
		return
	}
	//配置全局变量
	VPNGatewayCommonLibs.SetRegionId(myRegionId)
	fmt.Println("RegionId ", myRegionId)

	//载入全局的zk等信息,from conf/global/*/env.json
	if err = RegionConfig.LoadEnvironmentByRegion(); err != nil {
		fmt.Println("Load Region Config File fail,", err)
		return
	}
	RegionConfig.DumpConfigContent()


	dir, _ := cfg.GetConfigByKey("log.LogDir")
	prefix, _ := cfg.GetConfigByKey("log.LogPrefix")
	suffix, _ := cfg.GetConfigByKey("log.LogSuffix")
	log_size, _ := cfg.GetConfigByKey("log.LogSize")
	log_level, _ := cfg.GetConfigByKey("log.LogLevel")

	if err = ufcommon.LoadConfigFromFile(option); err != nil {
		fmt.Println("Load Config File fail,", err)
		return
	}
	ufcommon.DumpConfigContent()
	fmt.Println(cfg)
	fmt.Println(dir)
	fmt.Println(prefix)
	fmt.Println(suffix)
	fmt.Println(log_size)
	fmt.Println(log_level)

	//初始化日志
	//dir, _ := ufcommon.GetConfigByKey("log.LogDir")
	//prefix, _ := ufcommon.GetConfigByKey("log.LogPrefix")
	//suffix, _ := ufcommon.GetConfigByKey("log.LogSuffix")
	//log_size, _ := ufcommon.GetConfigByKey("log.LogSize")
	//log_level, _ := ufcommon.GetConfigByKey("log.LogLevel")
	//uflog.InitLogger(dir.(string), prefix.(string), suffix.(string), int64(log_size.(float64)), log_level.(string))
	//
	//regionId, err := ufcommon.GetOptionStringValue("r")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//myRegionId, err := strconv.Atoi(regionId)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//if myRegionId == 0 {
	//	fmt.Println("RegionId can not be zero ...")
	//	return
	//}
	//
	////配置全局变量
	//CNatCommon.SetRegionId(myRegionId)
	//fmt.Println("RegionId ", myRegionId)
	//
	////载入全局的zk等信息,from conf/global/*/env.json
	//if err = CNatCommon.LoadEnvironmentByRegion(); err != nil {
	//	fmt.Println("Load Config File fail,", err)
	//	return
	//}
	//CNatCommon.DumpConfigContent()
	vpngwmonitor.InitNameService()
	ufcommon.InitWrapPanic(true)
	err = vpngwmonitor.Prepare()
	//if err != nil {
	//	fmt.Println("Init error : ", err)
	//	return
	//}
	//// 启动定时任务
	//uftask.TimerTaskServe()
	//
	//var w sync.WaitGroup
	//w.Add(1)
	//w.Wait()
}
