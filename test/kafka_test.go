package test

import (
	"context"
	"dragon/core/dragon/conf"
	"dragon/tools/kafka"
	"fmt"
	"log"
	"testing"
)

// test produce
func TestProduce(t *testing.T) {
	err := kafka.Produce("test", "hello kafka")
	log.Println("kafka.Produce err", err)
}

// test Consumer
func TestConsume(t *testing.T) {
	// todo 本地业务处理时要注意记录offset,下次启动从offset开始
	conf.InitConf()
	r := kafka.GetConsumerConn("test", 25)
	for {
		m, err := r.ReadMessage(context.Background())
		log.Println(22222)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	r.Close()
}
