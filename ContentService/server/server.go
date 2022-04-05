package server

import (
	"github.com/shivamsouravjha/Micro-Game/ContentService/config"
	"github.com/shivamsouravjha/Micro-Game/ContentService/routes"
)

func Init() {
	r := routes.NewRouter()
	r.Run(":" + config.Get().PORT)
}
