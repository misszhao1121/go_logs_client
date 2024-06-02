package tailf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"time"
)

type CollectConf struct {
	LogPath string
	Topic   string
}

type MessAge struct {
	Msg   string
	Topic string
}
type TailObject struct {
	tail *tail.Tail
	conf CollectConf
}

type TailObjMgr struct {
	tailObjs []*TailObject
	msgChan  chan *MessAge
}

var (
	tailObjMgr *TailObjMgr
)

func GetLine() (msg *MessAge) {
	msg = <-tailObjMgr.msgChan
	return
}

func InitTail(conf []CollectConf, queueSize int) (err error) {
	if len(conf) == 0 {
		fmt.Errorf("inaviad config for logs collect,conf:%v", conf)
		return
	}
	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *MessAge, queueSize),
	}
	for _, v := range conf {
		obj := &TailObject{
			conf: v,
		}
		tails, errTail := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,
			Follow:    true,
			MustExist: false,
			Poll:      true,
		})
		if errTail != nil {
			err = errTail
			return
		}
		obj.tail = tails

		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

		go readFromTail(obj)
	}

	return
}
func readFromTail(object *TailObject) {
	for true {
		line, ok := <-object.tail.Lines
		if !ok {
			logs.Warning("tail file close repon,filename:%s\n", object.tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		messAge := &MessAge{
			Msg:   line.Text,
			Topic: object.conf.Topic,
		}
		tailObjMgr.msgChan <- messAge
	}
}
