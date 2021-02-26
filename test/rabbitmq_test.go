package test

import (
	"dragon/tools/rabbitmq"
	"github.com/go-dragon/util"
	"log"
	"testing"
)

// test publish
func TestPublish(t *testing.T) {
	mq, err := rabbitmq.New("amqp://guest:guest@localhost:5672/", "test_ex", "direct", "test_queue", "test_key")
	if err != nil {
		log.Fatal(err)
	}
	msg := "hello testgo " + util.RandomStr(4)
	log.Println("发送信息:", msg)
	err = mq.Publish(msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("消息发布结果:", (<-mq.ChanConfirm).Ack)
	mq.Close()
}

func TestManyPublish(t *testing.T) {
	mq, err := rabbitmq.New("amqp://guest:guest@localhost:5672/", "test_ex", "direct", "test_queue", "test_key")
	if err != nil {
		log.Fatal(err)
	}
	for n := 0; n < 100000; n++ {
		msg := "hello testgo " + util.RandomStr(10)
		log.Println("发送信息:", msg)
		err = mq.Publish(msg)
		if err != nil {
			log.Fatal(err)
		}
		confirm := <-mq.ChanConfirm
		log.Println("消息发布结果:", confirm.Ack)
	}
	mq.Close()
}
// 500/s
func BenchmarkPublish(b *testing.B) {
	mq, err := rabbitmq.New("amqp://guest:guest@localhost:5672/", "test_ex", "direct", "test_queue", "test_key")
	if err != nil {
		log.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		msg := "hello testgo " + util.RandomStr(10)
		log.Println("发送信息:", msg)
		err = mq.Publish(msg)
		if err != nil {
			log.Fatal(err)
		}
		confirm := <-mq.ChanConfirm
		log.Println("消息发布结果:", confirm.Ack)
	}
	mq.Close()
}

// test consumer 10k/s
func TestGetConsumer(t *testing.T) {
	mq, err := rabbitmq.New("amqp://guest:guest@localhost:5672/", "test_ex", "direct", "test_queue", "test_key")
	if err != nil {
		log.Fatal(err)
	}
	consumer, err := mq.GetConsumer("go-consumer")
	if err != nil {
		log.Fatal(err)
	}
	for msg := range consumer {
		log.Println("msg", string(msg.Body))
		msg.Ack(false)
	}
}
