package services

import (
	"Reddit-Clone/src/share/config"
	"Reddit-Clone/src/share/pkg/custome_logger"
	"Reddit-Clone/src/share/pkg/queue"
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQService struct {
	cfg     *config.Config
	channel *amqp.Channel
	cq      *amqp.Queue
}

func NewRabbitMQService(cfg *config.Config, lg custome_logger.Logger, queueName string, dto *queue.CreateQueueDto) *RabbitMQService {
	ch := queue.GetRabbitChanel()
	var createdQueue *amqp.Queue
	if dto == nil {
		dto = &queue.CreateQueueDto{Name: queueName}
	}
	createdQueue, err := queue.CreateRabbitQueue(ch, *dto)
	if err != nil {
		lg.Error(custome_logger.RabbitMq, custome_logger.CreateQueue, err.Error(), nil)
		return nil
	}
	return &RabbitMQService{
		cfg:     cfg,
		channel: ch,
		cq:      createdQueue,
	}
}

func (s *RabbitMQService) PublishMessage(dto queue.PublishMessageDto) error {
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

	return s.channel.Publish(
		dto.Exchange,
		s.cq.Name,
		dto.Mandatory,
		dto.Immediate,
		message,
	)
}
func (s *RabbitMQService) ConsumeMessage(dto queue.ConsumeMessageDto) (<-chan amqp.Delivery, error) {
	return s.channel.Consume(
		s.cq.Name,
		dto.Consumer,
		dto.AutoAck,
		dto.Exclusive,
		dto.NoLocal,
		dto.NoWait,
		dto.Args,
	)
}
