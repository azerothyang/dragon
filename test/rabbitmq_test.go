package test

import (
	"dragon/tools/rabbitmq"
	"github.com/go-dragon/util"
	"log"
	"testing"
)

// test Consumer
func TestPublish(t *testing.T) {
	mq := rabbitmq.New("amqp://guest:guest@127.0.0.1:5672/", "testgo")
	msg := "hello testgo " + util.RandomStr(4)
	log.Println("发送信息:", msg)
	ch , _ := mq.Publish(msg)
	log.Println("消息发布结果:", (<-ch).Ack)
	close(ch)
}

// test produce
func TestGetConsumer(t *testing.T) {
	mq := rabbitmq.New("amqp://guest:guest@127.0.0.1:5672/", "testgo")
	consumer, err := mq.GetConsumer("go-consumer")
	if err != nil {
		log.Fatal(err)
	}
	for msg := range consumer {
		log.Println("msg", string(msg.Body))
		msg.Ack(false)
	}
}
