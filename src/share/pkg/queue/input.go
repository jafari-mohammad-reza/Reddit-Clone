package queue

import "github.com/streadway/amqp"

type CreateQueueDto struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Arguments  amqp.Table
}

type PublishMessageDto struct {
	Exchange    string
	Mandatory   bool
	Immediate   bool
	ContentType string
	Body        interface{}
}

type ConsumeMessageDto struct {
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}
