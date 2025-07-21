package mq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MQClient struct {
	Conn         *amqp.Connection
	Channel      *amqp.Channel
	OrderQueue   *amqp.Queue
	RoutingKey   string
	ExchangeName string
}

func (mqc *MQClient) Close() {
	mqc.Channel.Close()
	mqc.Conn.Close()
}

func InitRabbitmqClient() *MQClient {
	const (
		ExchangeFoodOrder = "foodOrderSystem"
		QueueOrderCreate  = "order-service.create"
		RoutingKeyOrders  = "orders"
	)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panicf("Rabbitmq failed to connect err: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("Rabbitmq failed to create channel err: %v", err)
	}

	orderQueue, err := ch.QueueDeclare(
		QueueOrderCreate,
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
		ExchangeFoodOrder,
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
		QueueOrderCreate,
		RoutingKeyOrders,
		ExchangeFoodOrder,
		false,
		nil,
	)
	if err != nil {
		log.Panicf("Rabbitmq failed to bind Queue err: %v", err)
	}

	return &MQClient{
		Conn:         conn,
		Channel:      ch,
		OrderQueue:   &orderQueue,
		RoutingKey:   RoutingKeyOrders,
		ExchangeName: ExchangeFoodOrder,
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
