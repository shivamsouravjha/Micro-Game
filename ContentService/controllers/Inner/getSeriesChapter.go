package inner

import (
	rabbitMQ "github.com/shivamsouravjha/Micro-Game/ContentService/helpers/interconnect"
)

func GetSeriesChapter(UserId string) string {
	return rabbitMQ.NewUserContent(UserId)
}
