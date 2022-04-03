package main

import (
	"github.com/shivamsouravjha/Micro-Game/DailyPassService/server"
	"github.com/shivamsouravjha/Micro-Game/DailyPassService/services"
)

func main() {
	go services.Run()
	server.Init()
}
