package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

// Rabbit struct
type Rabbit struct {
	Dsn       string
	Conn      *amqp.Connection
	Channel   *amqp.Channel
	Queue     *amqp.Queue
	QueueName string
	Consumer  <-chan amqp.Delivery
}

// New a Rabbit
// dsn amqp://guest:guest@localhost:5672/
func New(dsn string, queueName string) *Rabbit {
	// connection
	conn, err := amqp.Dial(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	// queue
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	return &Rabbit{
		Dsn:       dsn,
		Conn:      conn,
		Channel:   ch,
		Queue:     &q,
		QueueName: queueName,
	}
}

// Close
func (r *Rabbit) Close() {
	r.Channel.Close()
	r.Conn.Close()
}

// Publish
func (r *Rabbit) Publish(body string) error {
	err := r.Channel.Publish("", r.QueueName, false, false, amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
		Body:         []byte(body),
	})
	return err
}

// Consumer, msg need to ack
func (r *Rabbit) GetConsumer(consumerName string) (<-chan amqp.Delivery, error) {
	return r.Channel.Consume(
		r.QueueName,  // queue
		consumerName, // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
}
