package main

import (
	"github.com/shivamsouravjha/Micro-Game/DailyPassService/helpers/rabbitMQ"
	"github.com/shivamsouravjha/Micro-Game/DailyPassService/server"
)

func main() {
	go rabbitMQ.Run()
	server.Init()
}
