package kafka

import (
	"github.com/IBM/sarama"
	"github.com/astaxie/beego/logs"
)

var (
	client sarama.SyncProducer
)

func InitKfka(addr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		logs.Error("init kafka failed,err:", err)
		return
	}
	logs.Debug("kafka conection success")
	return
}
func SendToKfka(data, topic string) (err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error("send kafka failed,pid:%v,offset:%v,topic:%v\n", pid, offset, topic)
		return
	}
	logs.Debug("send succ,pid:%v,offset:%v,topic:%v\n", pid, offset, topic)
	return
}
