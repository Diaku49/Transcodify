package mq

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MQClient struct {
	Conn             *amqp.Connection
	Channel          *amqp.Channel
	UploadVideoQueue *amqp.Queue
	RoutingKey       string
	ExchangeName     string
}

func (mqc *MQClient) Close() {
	mqc.Channel.Close()
	mqc.Conn.Close()
}

func InitRabbitmqClient() *MQClient {
	const (
		ExchangeVideoTranscoder = "videotranscode"
		QueueUploadVideo        = "upload-video-service"
		RoutingKeyUpload        = "upload-video"
	)

	url := os.Getenv("RMQ_URL")
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Panicf("Failed connecting to rabbitmq err: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("Failed creating channel err: %v", err)
	}

	uploadVideoQueue, err := ch.QueueDeclare(
		QueueUploadVideo,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panicf("Rabbitmq failed to create the queue err: %v", err)
	}

	err = ch.ExchangeDeclare(
		ExchangeVideoTranscoder,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panicf("Rabbitmq failed to create the exchange err: %v", err)
	}

	err = ch.QueueBind(
		QueueUploadVideo,
		RoutingKeyUpload,
		ExchangeVideoTranscoder,
		false,
		nil,
	)
	if err != nil {
		log.Panicf("Rabbitmq failed to bind Queue err: %v", err)
	}

	return &MQClient{
		Conn:             conn,
		Channel:          ch,
		UploadVideoQueue: &uploadVideoQueue,
		RoutingKey:       RoutingKeyUpload,
		ExchangeName:     ExchangeVideoTranscoder,
	}
}

func (rmqp *MQClient) MakeInfoQueue(id string) error {
	_, err := rmqp.Channel.QueueDeclare(
		"video-info-"+id,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-expires": int32(300 * 60 * 1000),
		},
	)
	if err != nil {
		return fmt.Errorf("couldnt make the info queue: %v", err)
	}

	err = rmqp.Channel.QueueBind(
		"video-info-"+id,
		id,
		rmqp.ExchangeName,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("rabbitmq failed to bind Queue err: %v", err)
	}

	return nil
}
