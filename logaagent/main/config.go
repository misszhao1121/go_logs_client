package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"logs-client/tail"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel    string
	logPath     string
	queueSize   int
	kafkaAddr   string
	CollectConf []tailf.CollectConf
}

//type CollectConf struct {
//	LogPath string
//	Topic   string
//}

func loadCollectConf(conf config.Configer) (err error) {
	var cc tailf.CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect::topic")
		return
	}
	appConfig.CollectConf = append(appConfig.CollectConf, cc)
	return
}

func loadConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed,err:", err)
		return
	}
	appConfig = &Config{}
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}
	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "./logs"
	}
	appConfig.queueSize, err = conf.Int("collect::queue_size")
	if err != nil {
		appConfig.queueSize = 200
	}
	appConfig.kafkaAddr = conf.String("kafka::server_addr")
	if len(appConfig.kafkaAddr) == 0 {
		//appConfig.kafkaAddr = "127.0.0.1:9092"
		err = fmt.Errorf("error kafka server,err", err)
		return
	}
	err = loadCollectConf(conf)
	if err != nil {
		fmt.Printf("load collect conf failed,err:%v\n")
		return
	}
	return
}
