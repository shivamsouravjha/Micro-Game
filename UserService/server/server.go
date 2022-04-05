package server

import "github.com/shivamsouravjha/Micro-Game/UserService/routes"

func Init() {
	r := routes.NewRouter()
	r.Run(":" + "8001")
}
