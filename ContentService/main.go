package main

import (
	"github.com/shivamsouravjha/Micro-Game/ContentService/helpers/rabbitMQ"
	"github.com/shivamsouravjha/Micro-Game/ContentService/server"
)

func main() {
	go rabbitMQ.Run("CreateUser")
	server.Init()
}
