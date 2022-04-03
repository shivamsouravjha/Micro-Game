package services

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Service -
type Service interface {
	Connect() error
	Publish(message string) error
}

// RabbitMQ -
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// Connect - establishes a connection to our RabbitMQ instance
// and declares the queue we are going to be using
func (r *RabbitMQ) Connect() error {
	fmt.Println("Connecting to RabbitMQ")
	var err error
	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Successfully Connected to RabbitMQ")

	// We need to open a channel over our AMQP connection
	// This will allow us to declare queues and subsequently consume/publish
	// messages
	r.Channel, err = r.Conn.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Here we declare our new queue that we want to publish to and consume
	// from:
	_, err = r.Channel.QueueDeclare(
		"CreateUser", // Queue Name
		false,        // durable
		false,        // Delete when not used
		false,        // exclusive
		false,        // no wait
		nil,          // additional args
	)
	return nil
}

// Publish - publishes a message to the queue
func (r *RabbitMQ) Publish(message string) error {
	// attempt to publish a message to the queue!
	err := r.Channel.Publish(
		"",
		"CreateUser",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// NewRabbitMQService - returns a pointer to a new RabbitMQ service
func NewRabbitMQService() *RabbitMQ {
	return &RabbitMQ{}
}

type App struct {
	Rmq *RabbitMQ
}

func Run(message string) error {
	rmq := NewRabbitMQService()
	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		return err
	}
	defer app.Rmq.Conn.Close()
	err = app.Rmq.Publish(message)

	fmt.Println("Successfull published message")
	if err != nil {
		return err
	}

	return nil
}
