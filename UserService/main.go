package main

import (
	"fmt"

	"github.com/shivamsouravjha/Micro-Game/UserService/server"
	"github.com/shivamsouravjha/Micro-Game/UserService/services"
)

type App struct {
	Rmq *services.RabbitMQ
}

func Run() error {
	fmt.Println("Go RabbitMQ Crash Course")

	rmq := services.NewRabbitMQService()
	app := App{
		Rmq: rmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		return err
	}
	defer app.Rmq.Conn.Close()

	err = app.Rmq.Publish("Hi")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	Run()
	server.Init()
}
