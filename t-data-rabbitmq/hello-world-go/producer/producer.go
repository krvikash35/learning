package producer

import (
	"hello-rmq/config"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer interface {
	Start()
	Stop()
	Send(queueName, msg string)
}

type producer struct {
	rmqConfig     config.Rabbitmq
	rmqConnection *amqp.Connection
	rmqChannel    *amqp.Channel
}

func NewProducer(config config.Rabbitmq) Producer {
	return &producer{
		rmqConfig: config,
	}
}

func (p *producer) Start() {
	url := p.rmqConfig.GetDialURL()
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Panicf("failed to create tcp connectiion to rabbitmq: %s", url)
	}
	p.rmqConnection = conn

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("failed to create channel to rabbitmq: %s", url)
	}
	p.rmqChannel = ch
}

func (p *producer) Stop() {

}

func (p *producer) Send(queueName, msg string) {
	q, err := p.rmqChannel.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Panicf("failed to declare queue: %s", err)
	}
	err = p.rmqChannel.Publish("", q.Name, false, false, amqp.Publishing{ContentType: "text/plan", Body: []byte(msg)})
	if err != nil {
		log.Panicf("failed to publish msg: %s", err)
	}
}
