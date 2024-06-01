package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"logs-client/tail"
)

func main() {
	fileName := "./config/logagent.conf"
	err := loadConf("ini", fileName)
	if err != nil {
		fmt.Printf("load conf failed, err:%v\n", err)
		panic("load conf failed")
		return
	}

	err = loadLogger()
	if err != nil {
		fmt.Printf("load logger failed, err:%v\n", err)
		panic("load conf failed")
		return
	}

	logs.Debug("load conf success,config:%v", appConfig)

	err = tail.InitTail(appConfig.CollectConf)
	if err != nil {
		logs.Error("init tail failed,err:%v", err)
		return
	}
	logs.Debug("initalIize all success")
	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed, err:%v", err)
		return
	}
	logs.Info("program exited")
}
