package main

import (
	"hello-rmq/config"
	"hello-rmq/producer"
)

func main() {

	c := config.Rabbitmq{
		Host:     "localhost",
		Port:     "5672",
		UserName: "guest",
		Password: "guest",
	}
	p1 := producer.NewProducer(c)
	p1.Start()
	p1.Send("test", "hello")
}
