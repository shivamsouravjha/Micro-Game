package server

import "github.com/shivamsouravjha/Micro-Game/routes"

func Init() {
	r := routes.NewRouter()
	r.Run(":" + "4000")
}
