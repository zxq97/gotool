package kafka

import (
	"context"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gotool/concurrent"
)

type Consumer struct {
	apiLogger *log.Logger
	excLogger *log.Logger
	consumer  *cluster.Consumer
	done      chan struct{}
	group     string
}

func InitConsumer(broker, topics []string, group string, apiLogger, excLogger *log.Logger) (*Consumer, error) {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Group.Return.Notifications = true
	consumer, err := cluster.NewConsumer(broker, group, topics, config)
	if err != nil {
		return nil, err
	}
	done := make(chan struct{})
	return &Consumer{consumer: consumer, apiLogger: apiLogger, excLogger: excLogger, done: done, group: group}, nil
}

func (consumer *Consumer) Start(fn func(ctx context.Context, msg *KafkaMessage)) {
	concurrent.Go(func() {
		for {
			select {
			case msg, ok := <-consumer.consumer.Messages():
				if ok {
					kfkmsg, err := unmarshal(msg.Value)
					if err != nil {
						consumer.excLogger.Println("consumer group unmarshal err", consumer.group, string(msg.Value), err)
						continue
					}
					ctx, cancel := consumerContext(kfkmsg.TraceId)
					now := time.Now()
					fn(ctx, kfkmsg)
					consumer.consumer.MarkOffset(msg, "")
					cancel()
					consumer.apiLogger.Println("consumer group since", consumer.group, time.Since(now))
				}
			case err := <-consumer.consumer.Errors():
				consumer.excLogger.Println("consumer group errors", consumer.group, err)
			case nft := <-consumer.consumer.Notifications():
				consumer.apiLogger.Println("consumer group notifications", consumer.group, nft)
			case <-consumer.done:
				consumer.apiLogger.Println("consumer group done", consumer.group)
				return
			}
		}
	})
}

func (consumer *Consumer) Stop() error {
	close(consumer.done)
	consumer.apiLogger.Println("consumer group stop", consumer.group)
	return consumer.consumer.Close()
}

func unmarshal(message []byte) (*KafkaMessage, error) {
	kfkmsg := &KafkaMessage{}
	err := proto.Unmarshal(message, kfkmsg)
	return kfkmsg, err
}
