package rabbitMQ

import (
	"fmt"

	Inner "github.com/shivamsouravjha/Micro-Game/ContentService/controllers/Inner"
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
func (r *RabbitMQ) Connect(channel string) error {
	fmt.Println("Connecting to RabbitMQ")
	var err error
	r.Conn, err = amqp.Dial("amqp://guest:guest@rabbitmq/")
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
		channel, // Queue Name
		false,   // durable
		false,   // Delete when not used
		false,   // exclusive
		false,   // no wait
		nil,     // additional args
	)
	return nil
}

func (r *RabbitMQ) Publish(channel string, message string) error {
	// attempt to publish a message to the queue!
	err := r.Channel.Publish(
		"",
		channel,
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

func Run(channel string) {
	fmt.Println("Go RabbitMQ Tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}
	msgs, _ := ch.Consume(
		channel,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
			UserId := string(d.Body)
			UserSeries := Inner.GetSeriesChapter(UserId)
			RunPublish("UserSeries", UserSeries)
			fmt.Println(UserSeries)
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}

func RunPublish(channel string, message string) error {
	rmq := NewRabbitMQService()
	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect(channel)
	if err != nil {
		return err
	}
	defer app.Rmq.Conn.Close()
	err = app.Rmq.Publish(channel, message)

	if err != nil {
		return err
	}
	fmt.Println("Successfull published message", message)

	return nil
}
