package main

import (
	"github.com/astaxie/beego/logs"
	"logs-client/logaagent/kafka"
	"logs-client/tail"
	"time"
)

func serverRun() (err error) {
	for {
		msg := tailf.GetLine()
		err := sendToKfka(msg)
		if err != nil {
			logs.Error("send to kafka err1:%v", err)
			time.Sleep(time.Second)
			continue
		}
	}
	return
}

func sendToKfka(msg *tailf.MessAge) (err error) {
	err = kafka.SendToKfka(msg.Msg, msg.Topic)
	return
}
