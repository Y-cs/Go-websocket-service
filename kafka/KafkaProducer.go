package kafka

import (
	"errors"
	"github.com/Shopify/sarama"
	"fmt"
)

type Producer struct {
	producer sarama.SyncProducer
}

func (producer *Producer) Init(ip string,port int) error {
	//这里可以初始化多个kafka的,因为是集群，最好多传几个，但是只传一个也可以使用
	servers := []string{fmt.Sprintf("%s:%d", ip, port)}
	p, err := sarama.NewSyncProducer(servers, sarama.NewConfig())
	if err != nil {
		return err
	}
	producer.producer = p
	return nil
}

func (producer *Producer) SendMessage(topic string, data []byte) error {
	if producer.producer == nil {
		return errors.New("no producer while send message")
	}
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   nil,
		Value: sarama.ByteEncoder(data),
	}
	_, _, err := producer.producer.SendMessage(kafkaMsg)
	return err
}

func (producer *Producer) close() error {
	if producer.producer != nil {
		return producer.producer.Close()
	}
	return nil
}