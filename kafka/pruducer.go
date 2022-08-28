package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

type Producer struct {
	apiLogger *log.Logger
	excLogger *log.Logger
	producer  sarama.SyncProducer
}

func InitKafkaProducer(addr []string, apiLogger, excLogger *log.Logger) (*Producer, error) {
	kfkConf := sarama.NewConfig()
	kfkConf.Producer.RequiredAcks = sarama.WaitForAll
	kfkConf.Producer.Retry.Max = 3
	kfkConf.Producer.Return.Successes = true
	kfkConf.Net.DialTimeout = defaultDialTimeout
	kfkConf.Net.ReadTimeout = defaultReadTimeout
	kfkConf.Net.WriteTimeout = defaultWriteTimeout
	producer, err := sarama.NewSyncProducer(addr, kfkConf)
	if err != nil {
		excLogger.Println("InitKafkaProducer err", addr, err)
		return nil, err
	}
	return &Producer{producer: producer, apiLogger: apiLogger}, nil
}

func (producer *Producer) SendMessage(topic string, key []byte, data []byte) error {
	partition, offset, err := producer.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(data),
	})
	if err != nil {
		producer.excLogger.Println("SendMessage topic date", topic, string(data), err)
		return err
	}
	producer.apiLogger.Println("SendMessage partition offset date", partition, offset, string(data))
	return nil
}

func (producer *Producer) Stop() error {
	producer.apiLogger.Println("producer stop")
	return producer.producer.Close()
}
