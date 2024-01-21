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
	connectionUrl := fmt.Sprintf("amqp://%s:%s@%s:%s", cfg.RabbitMq.User, cfg.RabbitMq.Password, cfg.RabbitMq.Host, cfg.RabbitMq.Port)
	rabbiMqConnection, err = amqp.Dial(connectionUrl)
	if err != nil {
		lg.Error(custome_logger.RabbitMq, custome_logger.Connect, err.Error(), nil)

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
