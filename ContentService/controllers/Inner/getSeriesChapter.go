package inner

import "github.com/shivamsouravjha/Micro-Game/ContentService/helpers/db"

func GetSeriesChapter(UserId string) string {
	return db.NewUserContent(UserId)
}
