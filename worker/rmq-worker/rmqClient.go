package rmqworker

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type MQClient struct {
	DB               *gorm.DB
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

func InitMQClient(db *gorm.DB) *MQClient {
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
		DB:               db,
		Conn:             conn,
		Channel:          ch,
		UploadVideoQueue: &uploadVideoQueue,
		RoutingKey:       RoutingKeyUpload,
		ExchangeName:     ExchangeVideoTranscoder,
	}
}

func (rmqp *MQClient) PublishToInfoQueue(id, msg string) {
	rmqp.Channel.Publish(
		rmqp.ExchangeName,
		id,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}
