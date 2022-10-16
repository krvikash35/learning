package config

import "fmt"

type Rabbitmq struct {
	Host     string
	Port     string
	UserName string
	Password string
}

func (rc *Rabbitmq) GetDialURL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", rc.UserName, rc.Password, rc.Host, rc.Port)
}
