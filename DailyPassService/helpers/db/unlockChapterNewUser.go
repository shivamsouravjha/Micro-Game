package db

import (
	"encoding/json"
	"fmt"

	structss "github.com/shivamsouravjha/Micro-Game/DailyPassService/services"
	structs "github.com/shivamsouravjha/Micro-Game/DailyPassService/struct"
)

func UnlockUserContentDAO(userSeries []byte) {
	var UserContent structs.NewUserContent
	json.Unmarshal(userSeries, &UserContent)
	unlockedChapters, _ := json.Marshal(UserContent.UnlockedChapters)
	_, err := structss.Dbmap.Exec("INSERT INTO dailypass (userId,unlockedChapters) VALUES(?,?)", UserContent.UserId, string(unlockedChapters))

	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
