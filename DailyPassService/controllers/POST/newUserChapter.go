package post

import (
	"github.com/shivamsouravjha/Micro-Game/DailyPassService/helpers/db"
)

func UnlockContentNewUser(userId string) {
	go db.UnlockUserContentDAO(userId)
}
