package post

import (
	"github.com/shivamsouravjha/Micro-Game/DailyPassService/helpers/db"
)

func UnlockContentNewUser(userSeries []byte) {
	go db.UnlockUserContentDAO(userSeries)
}
