package server

import (
	"github.com/shivamsouravjha/Micro-Game/UserService/config"
	"github.com/shivamsouravjha/Micro-Game/UserService/routes"
)

func Init() {
	r := routes.NewRouter()
	r.Run(":" + config.Get().PORT)
}
