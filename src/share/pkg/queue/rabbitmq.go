package queue

import (
	"fmt"
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
	"github.com/streadway/amqp"
)

var rabbiMqConnection *amqp.Connection
var rabbiMqChannel *amqp.Channel

func InitRabbitMq(cfg *config.Config, lg custome_logger.Logger) {
	var err error
	rabbiMqConnection, err = amqp.Dial("amqp://myuser:mypassword@localhost:5672/")
	if err != nil {
		lg.Error(custome_logger.RabbitMq, custome_logger.Connect, err.Error(), nil)

	}
	if rabbiMqConnection.ConnectionState().HandshakeComplete != true {
		msg := "failed to establish rabbitmq connection"
		lg.Error(custome_logger.RabbitMq, custome_logger.Connect, msg, nil)

	}
	rabbiMqChannel, err = rabbiMqConnection.Channel()
	if err != nil {
		lg.Error(custome_logger.RabbitMq, custome_logger.CreateChannel, err.Error(), nil)
	}
}

func GetRabbitConnection() *amqp.Connection {
	return rabbiMqConnection
}
func GetRabbitChanel() *amqp.Channel {
	return rabbiMqChannel
}

func CloseRabbitConnection(lg custome_logger.Logger) {
	err := rabbiMqConnection.Close()
	if err != nil {
		lg.Error(custome_logger.RabbitMq, custome_logger.Close, err.Error(), nil)
		panic(err)
	}
}

func CloseRabbitChanel(lg custome_logger.Logger) {
	err := rabbiMqChannel.Close()
	if err != nil {
		lg.Error(custome_logger.RabbitMq, custome_logger.Close, err.Error(), nil)
		panic(err)
	}
}

func CreateRabbitQueue(ch *amqp.Channel, dto CreateQueueDto) (*amqp.Queue, error) {
	queue, err := ch.QueueDeclare(
		dto.Name,
		dto.Durable,
		dto.AutoDelete,
		dto.Exclusive,
		dto.NoWait,
		dto.Arguments,
	)
	if err != nil {
		return nil, err
	}
	return &queue, nil
}

func PublishMessage(ch *amqp.Channel, queueName string, dto PublishMessageDto) error {
	var channel *amqp.Channel
	if ch != nil {
		channel = ch
	} else {
		channel = GetRabbitChanel()
	}
	contentType := "text/plain"
	if dto.ContentType != "" {
		contentType = dto.ContentType
	}

	body, ok := dto.Body.(string)
	if !ok {
		return fmt.Errorf("body must be a string")
	}

	message := amqp.Publishing{
		ContentType: contentType,
		Body:        []byte(body),
	}

	return channel.Publish(
		dto.Exchange,
		queueName,
		dto.Mandatory,
		dto.Immediate,
		message,
	)
}
func ConsumeMessage(ch *amqp.Channel, queueName string, dto ConsumeMessageDto) (<-chan amqp.Delivery, error) {
	var channel *amqp.Channel
	if ch != nil {
		channel = ch
	} else {
		channel = GetRabbitChanel()
	}
	return channel.Consume(
		queueName,
		dto.Consumer,
		dto.AutoAck,
		dto.Exclusive,
		dto.NoLocal,
		dto.NoWait,
		dto.Args,
	)
}
