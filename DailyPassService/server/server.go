package server

import "github.com/shivamsouravjha/Micro-Game/DailyPassService/routes"

func Init() {
	r := routes.NewRouter()
	r.Run(":" + "8003")
}
