package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"logs-client/logaagent/kafka"
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

	err = tailf.InitTail(appConfig.CollectConf, appConfig.queueSize)
	if err != nil {
		logs.Error("init tail failed,err:%v", err)
		return
	}

	err = kafka.InitKfka(appConfig.kafkaAddr)
	if err != nil {
		logs.Error("init kafka failed,kafkaAddr:%v", appConfig.kafkaAddr)
		return
	}
	logs.Debug("initialLi all success")

	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed, err:%v", err)
		return
	}
	logs.Info("program exited")
}
