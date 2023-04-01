package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main()  {
	//生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	//连接kafka
	client,err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"},config)
	if err != nil {
		fmt.Println("NewSyncProducer failed ,err :",err)
		return
	}
	defer client.Close()

	//封装消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "shopping"
	msg.Value = sarama.StringEncoder("top goer Neal is best!")

	//发送消息
	partition, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed ,err :",err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n",partition,offset)
}
