package queue

import (
	"github.com/streadway/amqp"
)

type Queue interface {
	Publish(queueName string, body []byte) error
}

type rabbitMQ struct {
	channel *amqp.Channel
}

func NewRabbitMQ(url string) (Queue, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &rabbitMQ{ch}, nil
}

func (r *rabbitMQ) Publish(queueName string, body []byte) error {
	_, err := r.channel.QueueDeclare(
		queueName, true, false, false, false, nil,
	)
	if err != nil {
		return err
	}
	return r.channel.Publish(
		"", queueName, false, false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         body,
		},
	)
}
